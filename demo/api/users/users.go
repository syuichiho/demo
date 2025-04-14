// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package users

import (
	"context"

	"demo/api/users/v1"
)

type IUsersV1 interface {
	Zhuce(ctx context.Context, req *v1.ZhuceReq) (res *v1.ZhuceRes, err error)
}
