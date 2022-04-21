package v1

type OsInfo struct {
	OsIp    string
	OsBasic []*InfoItem
}
type InfoItem struct {
	ItemName   string
	ItemCmd    string
	ItemStatus bool
	ItemJud    string
	ItemResult string
}
type OsInfoGroup struct {
	OsIfG []*OsInfo
}

var (
	osiplist = []string{"10.10.26.69", "10.10.26.70", "10.10.26.71"}
)

func NewOsInfoGroup() *OsInfoGroup {
	osGroup := make([]*OsInfo, len(osiplist))
	for i, osip := range osiplist {
		osGroup[i] = newOsInfo(osip)
	}
	return &OsInfoGroup{
		OsIfG: osGroup,
	}
}
func newOsInfo(osIp string) *OsInfo {

	//infoitem := make([]*InfoItem,len(osiplist))
	osBasic := []*InfoItem{
		&InfoItem{"diskspace", " df -h /data  |tail -1", true, "", ""},
		&InfoItem{"dns", " cat /etc/resolv.conf |head -1 ", false, "127.0.0.1", ""},
		&InfoItem{"firewalldstatus", " systemctl status firewalld.service |grep dead ", false, "dead", ""},
		&InfoItem{"selinuxstatus", " getenforce ", false, "Disabled", ""}}

	return &OsInfo{
		OsIp:    osIp,
		OsBasic: osBasic,
	}
}
