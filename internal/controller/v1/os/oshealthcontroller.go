package os

import srvv1 "bluekinghealth/internal/service/v1"
import "bluekinghealth/internal/store"

type OsHealthController struct {
	srv srvv1.Service
}

func NewOsHealthController(cmd store.StoreFactory) *OsHealthController {
	return &OsHealthController{
		srv: srvv1.NewService(cmd),
	}
}
