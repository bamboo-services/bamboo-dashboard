// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package initial

import (
	"context"

	"bamboo-dashboard/api/initial/v1"
)

type IInitialV1 interface {
	InitialConsoleUser(ctx context.Context, req *v1.InitialConsoleUserReq) (res *v1.InitialConsoleUserRes, err error)
}
