package store

import (
	"context"
)

//存储层子接口,这里是check子接口
type CheckCmd interface {
	PaasCheckCmd(ctx context.Context)([]string ,error)
	CmdbCheckCmd(ctx context.Context)([]string,error)
}
