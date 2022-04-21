package store

import (
	metav1 "bluekinghealth/pkg/meta/v1"
	"context"
)

type SaasCmd interface {
	SaasStatus(ctx context.Context, sg *metav1.SaasGroup) (string, error)
}
