// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package users

import (
	"context"

	"demo/api/users/v1"
)

type IUsersV1 interface {
	Denglu(ctx context.Context, req *v1.DengluReq) (res *v1.DengluRes, err error)
}
