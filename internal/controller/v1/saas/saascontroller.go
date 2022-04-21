package saas

import (
	srvv1 "bluekinghealth/internal/service/v1"
	"bluekinghealth/internal/store"
)

type SaasController struct {
	srv srvv1.Service
}

func NewSaasController(component store.StoreFactory) *SaasController {
	return &SaasController{srv: srvv1.NewService(component)}
}
