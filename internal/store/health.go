package store

import (
	metav1 "bluekinghealth/pkg/meta/v1"
	"context"
)

//存储层子接口,这里是check子接口
type HealthCmd interface {
	PaasCheckCmd(ctx context.Context, app []*metav1.App) (string, error)
	CmdbCheckCmd(ctx context.Context) ([]byte, error)
}
