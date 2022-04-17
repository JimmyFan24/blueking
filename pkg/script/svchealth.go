package script

const CheckScript = `
#!/usr/bin/env bash
# Description: check service's health wrapper

set -euo pipefail
SELF_DIR=$(dirname "$(readlink -f "$0")")

# 加载load_env和通用函数
source "${SELF_DIR}"/tools.sh

BK_PKG_SRC_PATH=${BK_PKG_SRC_PATH:-/data/src}
BK_HOME=${BK_HOME:-/data/bkee}
PCMD=${SELF_DIR}/pcmd.sh

SUPPORT_MODULE=(bkssm bkiam usermgr paas cmdb gse job consul bklog dbcheck bkmonitorv3)
usage () {
    echo "$0 <module>"
    echo "<module> can be one of the following:"
    echo "${SUPPORT_MODULE[@]}"
    exit 1
}
[[ $# -ne 1 ]] && usage

MODULE=$1
case $MODULE in
    bkssm|bkiam|usermgr) $PCMD -m "${MODULE#bk}" "$SELF_DIR/health_check/check_consul_svc_health.sh $MODULE" ;;
    paas|cmdb) $PCMD -m "$MODULE" "$SELF_DIR/health_check/check_consul_svc_health.sh ^${MODULE}-" ;;
    gse) $PCMD -m gse '$CTRL_DIR/health_check/check_gse.sh' ;;
    job) 
        step "check job backend health"
        $PCMD -m job '$CTRL_DIR/health_check/check_job.sh'
        step "check job frontend resource"
        $PCMD -m nginx 'runuser -u blueking -- ls -l $BK_HOME/job/frontend/index.html'
        ;;
    consul) $PCMD -m all '$CTRL_DIR/health_check/check_consul_resolv.sh' ;;
    bklog) $PCMD -m ${MODULE#bk} "$SELF_DIR/health_check/check_consul_svc_health.sh $MODULE" ;;
    bkmonitorv3|monitorv3) $PCMD -H $BK_MONITORV3_MONITOR_IP 'workon bkmonitorv3-monitor; ./bin/manage.sh healthz' ;;
    nginx) $PCMD -m nginx '$CTRL_DIR/health_check/check_openresty.sh' ;;
    dbcheck)
        set +u
        [[ -f ${HOME}/.bkrc ]] && source ${HOME}/.bkrc
        workon deploy_check && python ${SELF_DIR}/health_check/deploy_check.py  -d "${SELF_DIR}/bin/04-final" -m "cmdb,paas,ssm,bkiam,usermgr,gse,job,bknodeman,bkmonitorv3"
        ;;
    --list)
        printf "%s\n" "${SUPPORT_MODULE[@]}"
        ;;
    *) usage ;;
esac

`
const Check_test = `
#!/usr/bin/env bash
echo $1`
const StatusScript = `
#!/usr/bin/env bash
# 组件运行状态查看脚本 只能中控机执行
SELF_DIR=$(dirname "$(readlink -f "$0")")
source "${SELF_DIR}"/functions
source "${SELF_DIR}"/tools.sh
# set -e

module=$1
project=$2
# 支持判断一下target 
target=$3


if ! [ -z ${target} ]; then
    if ! grep "[0-9]" <<<${target} >/dev/null; then
        err "不支持多project"
    else
        module_ip=BK_${module^^}_IP${target}
        if [[ -z ${!module_ip} ]]; then
            err "${module_ip} 不存在"
        else
            target=${!module_ip}
        fi
    fi
else    
    if ! [ -z ${project} ]; then
        if grep "[0-9]" <<<${project} >/dev/null; then
            module_ip=BK_${module^^}_IP${project}
            if [[ -z ${!module_ip} ]]; then
                err "${module_ip} 不存在"
            else
                target=${!module_ip}
                project=""
            fi
        else
            target=${module}
        fi
    else
        target=${module}
    fi

fi

declare -a THIRD_PARTY_SVC=(
    consul
    consul-template
    mysql@[a-z]+
    redis@[a-z]+
    openresty
    rabbitmq-server
    zookeeper
    mongod
    kafka
    elasticsearch
    influxdb
    beanstalkd
)
TMP_PTN=$(printf "%s|" "${THIRD_PARTY_SVC[@]}")
THIRD_PARTY_SVC_PTN="^(${TMP_PTN%|})\.service$"

declare -A SERVICE=(
    ["mysql"]=mysql@default
    ['redis']=redis@default
    ["es7"]=elasticsearch
    ["nodeman"]=bk-nodeman
    ["consul"]=consul
    ["kafka"]=kafka
    ["usermgr"]=bk-usermgr
    ["redis_sentinel"]=redis-sentinel@default
    ["rabbitmq"]=rabbitmq-server
    ["zk"]=zookeeper
    ["mongodb"]=mongod
    ["influxdb"]=influxdb
    ["nginx"]=openresty
    ["beanstalk"]=beanstalkd
    ["yum"]=bk-yum
    ["fta"]=bk-fta
    ["iam"]=bk-iam
    ["ssm"]=bk-ssm
    ["license"]=bk-license
    ["appo"]=bk-paasagent
    ["appt"]=bk-paasagent
    ["nfs"]=nfs-server
    ['consul-template']=consul-template
    ['lesscode']=bk-lesscode
)

case $module in 
    paas|cmdb|gse|job)
        module=${module#bk}
        target_name=$(map_module_name "${module}")
        source <(/opt/py36/bin/python ${SELF_DIR}/qq.py -p ${BK_PKG_SRC_PATH}/${target_name}/projects.yaml -P ${SELF_DIR}/bin/default/port.yaml)
        if [[ -z ${project} ]]; then
            projects=${_projects["${module}"]}
            pcmdrc "${target}" "get_common_bk_service_status ${module} ${projects[*]}"
        else
            if [[ ${module}  == 'paas' ]];then
                pcmdrc "${target}" "get_spic_bk_service_status ${module} ${project}"
            else 
                pcmdrc "${target}" "get_common_bk_service_status ${module} ${project}"
            fi
        fi
        ;;
    monitorv3|bkmonitorv3|log|bklog)
        module=${module#bk*}
        target_name=$(map_module_name "$module")
        source <(/opt/py36/bin/python ${SELF_DIR}/qq.py -p ${BK_PKG_SRC_PATH}/${target_name}/projects.yaml -P ${SELF_DIR}/bin/default/port.yaml)
        if [ -z "${project}" ];then
            for project in ${_projects[${module}]};do
                emphasize "status ${module} ${project} on host: ${_project_ip["${target_name},${project}"]}"
                if [[ "${module}" =~ "log" ]]; then
                    pcmdrc "${_project_ip["${target_name},${project}"]}" "get_service_status bk-${module}-${project}"
                else
                    pcmdrc "${_project_ip["${target_name},${project}"]}" "get_service_status bk-${project}"
                fi
            done
        else
            emphasize "status ${module} ${project} on host: ${_project_ip["${target_name},${project}"]}"
            if [[ "${module}" =~ "log" ]]; then
                pcmdrc "${_project_ip["${target_name},${project}"]}" "get_service_status bk-${module}-${project}"
            else
                pcmdrc "${_project_ip["${target_name},${project}"]}" "get_service_status bk-${project}"
            fi
        fi
        ;;
    nginx)  
        pcmdrc "${target}" "get_service_status ${SERVICE[$module]} ${SERVICE["consul-template"]}"
        ;;
    yum)
        # 中控机安装模块
        pcmdrc "$LAN_IP" "get_service_status ${SERVICE[$module]}"
        ;;
    bkiam|bkssm)
        target_name=${module#bk}
        pcmdrc "${target_name}" "get_service_status ${SERVICE[${target_name}]}"
        ;;
    bknodeman|nodeman)
        target_name=${module#bk}
        pcmdrc "${target_name}" "get_service_status ${SERVICE[${target_name}]} ${SERVICE["consul-template"]} ${SERVICE["nginx"]}"
        ;;
    paas_plugins|paas_plugin)
        pcmdrc "${BK_PAAS_IP0}" "get_service_status bk-paas-plugins-log-alert"
        pcmdrc paas "get_service_status bk-logstash-paas-app-log  bk-filebeat@paas_esb_api" 
        if ! [ -z "${BK_APPT_IP_COMMA}" ]; then
            pcmdrc appt "get_service_status bk-filebeat@celery  bk-filebeat@store bk-filebeat@django bk-filebeat@java bk-filebeat@uwsgi"
        fi
        pcmdrc appo "get_service_status bk-filebeat@celery  bk-filebeat@store bk-filebeat@django bk-filebeat@java bk-filebeat@uwsgi"
        ;;
    bkall)
        pcmdrc all "FORCE_TTY=1 $CTRL_DIR/bin/bks.sh ^bk-"
        ;;
    tpall)
        pcmdrc all "FORCE_TTY=1 $CTRL_DIR/bin/bks.sh \"$THIRD_PARTY_SVC_PTN\" "
        ;;
    all)
        echo "Status of all blueking components: "
        pcmdrc all "FORCE_TTY=1 $CTRL_DIR/bin/bks.sh ^bk-"
        echo 
        echo "Status of all third-party components: "
        pcmdrc all "FORCE_TTY=1 $CTRL_DIR/bin/bks.sh \"$THIRD_PARTY_SVC_PTN\" "
        ;;
    *)  
        if [[ -z ${SERVICE[$module]} ]]; then
            echo  "当前不支持 '${module}' 的状态检测."
            exit 1
        fi
        pcmdrc "${target}" "get_service_status ${SERVICE[$module]}"
        ;;
esac
`