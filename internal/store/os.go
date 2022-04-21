package store

import (
	metav1 "bluekinghealth/pkg/meta/v1"
	"context"
)

//store impl
type OsHealthCmd interface {
	Basic(ctx context.Context, osinfo *metav1.OsInfoGroup) (string, error)
}
