package platform

const status_shell = `
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
esac`

const checkdata = `[1] 16:38:09 [SUCCESS] 10.10.26.69
bkssm(http://10.10.26.69:5000/healthz)       : true
[2] 16:38:09 [SUCCESS] 10.10.26.74
bkssm(http://10.10.26.74:5	000/healthz)       : true
[1] 16:38:10 [SUCCESS] 10.10.26.69
bkiam(http://10.10.26.69:5001/healthz)       : true
[2] 16:38:10 [SUCCESS] 10.10.26.74
bkiam(http://10.10.26.74:5001/healthz)       : true
[1] 16:38:11 [SUCCESS] 10.10.26.69
usermgr(http://10.10.26.69:8009/healthz/)    : true
[2] 16:38:11 [SUCCESS] 10.10.26.73
usermgr(http://10.10.26.73:8009/healthz/)    : true
[1] 16:38:14 [SUCCESS] 10.10.26.73
paas-apigw(http://10.10.26.73:8005/api/healthz/): true
paas-appengine(http://10.10.26.73:8000/v1/healthz/): true
paas-esb(http://10.10.26.73:8002/healthz/)   : true
paas-login(http://10.10.26.73:8003/healthz/) : true
paas-paas(http://10.10.26.73:8001/healthz/)  : true
[2] 16:38:14 [SUCCESS] 10.10.26.71
paas-apigw(http://10.10.26.71:8005/api/healthz/): true
paas-appengine(http://10.10.26.71:8000/v1/healthz/): true
paas-esb(http://10.10.26.71:8002/healthz/)   : true
paas-login(http://10.10.26.71:8003/healthz/) : true
paas-paas(http://10.10.26.71:8001/healthz/)  : true
[1] 16:38:17 [FAILURE] 10.10.26.75 Exited with error code 14
cmdb-admin(http://10.10.26.75:9000/healthz)  : false Reason: connection refused
cmdb-api(http://10.10.26.75:9001/healthz)    : false Reason: connection refused
cmdb-auth(http://10.10.26.75:9002/healthz)   : false Reason: connection refused
cmdb-cache(http://10.10.26.75:9014/healthz)  : false Reason: connection refused
cmdb-cloud(http://10.10.26.75:9003/healthz)  : false Reason: connection refused
cmdb-core(http://10.10.26.75:9004/healthz)   : false Reason: connection refused
cmdb-datacollection(http://10.10.26.75:9005/healthz): false Reason: connection refused
cmdb-event(http://10.10.26.75:9006/healthz)  : false Reason: connection refused
cmdb-host(http://10.10.26.75:9007/healthz)   : false Reason: connection refused
cmdb-operation(http://10.10.26.75:9008/healthz): false Reason: connection refused
cmdb-proc(http://10.10.26.75:9009/healthz)   : false Reason: connection refused
cmdb-task(http://10.10.26.75:9011/healthz)   : false Reason: connection refused
cmdb-topo(http://10.10.26.75:9012/healthz)   : false Reason: connection refused
cmdb-web(http://10.10.26.75:9013/healthz)    : false Reason: connection refused
[2] 16:38:17 [FAILURE] 10.10.26.71 Exited with error code 14
cmdb-admin(http://10.10.26.71:9000/healthz)  : false Reason: connection refused
cmdb-api(http://10.10.26.71:9001/healthz)    : false Reason: connection refused
cmdb-auth(http://10.10.26.71:9002/healthz)   : false Reason: connection refused
cmdb-cache(http://10.10.26.71:9014/healthz)  : false Reason: connection refused
cmdb-cloud(http://10.10.26.71:9003/healthz)  : false Reason: connection refused
cmdb-core(http://10.10.26.71:9004/healthz)   : false Reason: connection refused
cmdb-datacollection(http://10.10.26.71:9005/healthz): false Reason: connection refused
cmdb-event(http://10.10.26.71:9006/healthz)  : false Reason: connection refused
cmdb-host(http://10.10.26.71:9007/healthz)   : false Reason: connection refused
cmdb-operation(http://10.10.26.71:9008/healthz): false Reason: connection refused
cmdb-proc(http://10.10.26.71:9009/healthz)   : false Reason: connection refused
cmdb-task(http://10.10.26.71:9011/healthz)   : false Reason: connection refused
cmdb-topo(http://10.10.26.71:9012/healthz)   : false Reason: connection refused
cmdb-web(http://10.10.26.71:9013/healthz)    : false Reason: connection refused
[1] 16:38:43 [SUCCESS] 10.10.26.72
bk-gse-alarm   : running
bk-gse-api     : running
bk-gse-btsvr   : running
bk-gse-data    : running
bk-gse-dba     : running
bk-gse-procmgr : running
bk-gse-syncdata: running
bk-gse-task    : running
[2] 16:38:43 [SUCCESS] 10.10.26.73
bk-gse-alarm   : running
bk-gse-api     : running
bk-gse-btsvr   : running
bk-gse-data    : running
bk-gse-dba     : running
bk-gse-procmgr : running
bk-gse-syncdata: running
bk-gse-task    : running

                      check job backend health
[1] 16:38:45 [SUCCESS] 10.10.26.75
job-execute    : true
job-backup     : true
job-logsvr     : true
job-crontab    : true
job-config     : true
job-analysis   : true
job-gateway-management: true
job-manage     : true
[2] 16:38:45 [SUCCESS] 10.10.26.71
job-execute    : true
job-backup     : true
job-logsvr     : true
job-crontab    : true
job-config     : true
job-analysis   : true
job-gateway-management: true
job-manage     : true

                     check job frontend resource
[1] 16:38:46 [SUCCESS] 10.10.26.70
-rw-r--r-- 1 blueking blueking 1795 Apr 13 11:31 /data/bkce/job/frontend/index.html
[2] 16:38:46 [SUCCESS] 10.10.26.73
-rw-r--r-- 1 blueking blueking 1795 Apr 13 11:31 /data/bkce/job/frontend/index.html
[1] 16:38:48 [SUCCESS] 10.10.26.72
check_consul_process [OK]
check_consul_listen_udp_53 [OK]
check_consul_listen_tcp_8500 [OK]
check_consul_warning_svc [OK]
check_consul_critical_svc [OK]
check_resolv_conf_127.0.0.1 [OK]
[2] 16:38:48 [FAILURE] 10.10.26.75 Exited with error code 1
check_consul_process [OK]
check_consul_listen_udp_53 [OK]
check_consul_listen_tcp_8500 [OK]
check_consul_warning_svc [OK]
以下服务consul显示为critical，请确认: cmdb-admin cmdb-api cmdb-auth cmdb-cache cmdb-cloud cmdb-core cmdb-datacollection cmdb-event cmdb-host cmdb-operation cmdb-proc cmdb-task cmdb-topo cmdb-web
check_resolv_conf_127.0.0.1 [OK]
Stderr: check_consul_critical_svc [FAIL]
[3] 16:38:48 [SUCCESS] 10.10.26.69
check_consul_process [OK]
check_consul_listen_udp_53 [OK]
check_consul_listen_tcp_8500 [OK]
check_consul_warning_svc [OK]
check_consul_critical_svc [OK]
check_resolv_conf_127.0.0.1 [OK]
[4] 16:38:48 [SUCCESS] 10.10.26.73
check_consul_process [OK]
check_consul_listen_udp_53 [OK]
check_consul_listen_tcp_8500 [OK]
check_consul_warning_svc [OK]
check_consul_critical_svc [OK]
check_resolv_conf_127.0.0.1 [OK]
[5] 16:38:48 [FAILURE] 10.10.26.71 Exited with error code 1
check_consul_process [OK]
check_consul_listen_udp_53 [OK]
check_consul_listen_tcp_8500 [OK]
check_consul_warning_svc [OK]
以下服务consul显示为critical，请确认: cmdb-admin cmdb-api cmdb-auth cmdb-cache cmdb-cloud cmdb-core cmdb-datacollection cmdb-event cmdb-host cmdb-operation cmdb-proc cmdb-task cmdb-topo cmdb-web
check_resolv_conf_127.0.0.1 [OK]
Stderr: check_consul_critical_svc [FAIL]
[6] 16:38:48 [SUCCESS] 10.10.26.74
check_consul_process [OK]
check_consul_listen_udp_53 [OK]
check_consul_listen_tcp_8500 [OK]
check_consul_warning_svc [OK]
check_consul_critical_svc [OK]
check_resolv_conf_127.0.0.1 [OK]
[7] 16:38:48 [SUCCESS] 10.10.26.70
check_consul_process [OK]
check_consul_listen_udp_53 [OK]
check_consul_listen_tcp_8500 [OK]
check_consul_warning_svc [OK]
check_consul_critical_svc [OK]
check_resolv_conf_127.0.0.1 [OK]`

const status = `
[1] 14:46:51 [SUCCESS] 10.10.26.72
[2] 14:46:52 [SUCCESS] 10.10.26.75
bk-cmdb-core.service            inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-operation.service       inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-auth.service            inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-event.service           inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-datacollection.service  inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-cloud.service           inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-web.service             inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-host.service            inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-proc.service            inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-task.service            inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-admin.service           inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-api.service             inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-cache.service           deactivating  (stop-sigterm) 1min 4s ago (cmdb_cacheservi)
bk-cmdb-topo.service            inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
[3] 14:46:52 [SUCCESS] 10.10.26.70
[4] 14:46:52 [SUCCESS] 10.10.26.69
[5] 14:46:52 [SUCCESS] 10.10.26.73
[6] 14:46:52 [SUCCESS] 10.10.26.74
[7] 14:46:53 [SUCCESS] 10.10.26.71
bk-cmdb-core.service                    inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-operation.service               inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-auth.service                    inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-event.service                   inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-datacollection.service          deactivating  (stop-sigterm) 1min 4s ago (cmdb_datacollec)
bk-cmdb-cloud.service                   inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-web.service                     inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-host.service                    inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-proc.service                    inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-task.service                    inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-admin.service                   inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-api.service                     inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-cache.service                   inactive      (dead) 59s ago (code=exited, status=0/SUCCESS)
bk-cmdb-topo.service                    inactive      (dead) 1min 0s ago (code=exited, status=0/SUCCESS)
[1] 14:46:53 [SUCCESS] 10.10.26.72
[2] 14:46:53 [SUCCESS] 10.10.26.75
[3] 14:46:53 [SUCCESS] 10.10.26.69
[4] 14:46:53 [SUCCESS] 10.10.26.73
[5] 14:46:54 [SUCCESS] 10.10.26.70
[6] 14:46:54 [SUCCESS] 10.10.26.71
[7] 14:46:54 [SUCCESS] 10.10.26.74
`
