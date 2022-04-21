package oshealth

import (
	"bluekinghealth/internal/store"
	cmd "bluekinghealth/pkg/command"
	metav1 "bluekinghealth/pkg/meta/v1"
	"context"
)

type osStore struct {
	name string
}

func (o osStore) Basic(ctx context.Context, osinfo *metav1.OsInfoGroup) (string, error) {
	err := cmd.OsHealth(osinfo)
	/*for _,osItem:= range osinfo.OsIfG {
		//logrus.Infof("infocmd is :%v",infocmd)

		if osItem.OsIp == cmd.LANIP {
			for _, infoitem := range osItem.OsBasic {

			data, err := cmd.ExecCmd(infoitem.ItemCmd)
			if err != nil {
				return "", err
			}
			//infoitem.ItemResult = data
			if infoitem.ItemName == "diskspace" {
				diskReg := regexp.MustCompile(`\d+`)
				infoitem.ItemResult = diskReg.FindStringSubmatch(data)[0]
				//logrus.Infof("infoitem.ItemResult is %v",infoitem.ItemResult )
				diskmeg, _ := strconv.Atoi(infoitem.ItemResult)
				//logrus.Info(diskmeg)
				if diskmeg > 90 {
					infoitem.ItemStatus = false
				} else {
					//logrus.Info("set infoitem itemstatus true..")
					infoitem.ItemStatus = true
				}
			} else if strings.Contains(data, infoitem.ItemJud) {
				infoitem.ItemStatus = true
				infoitem.ItemResult = data
			} else {
				infoitem.ItemStatus = false
			}

		}
		}else {

		}
	}*/
	if err != nil {
		return "", err
	}
	return "", nil
}

var _ store.OsHealthCmd = osStore{}

func newOsHealthStore(name string) store.OsHealthCmd {
	return &osStore{
		name: name,
	}
}
