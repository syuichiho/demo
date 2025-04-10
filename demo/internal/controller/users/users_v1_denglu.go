package users

import (
	"context"

	v1 "demo/api/users/v1"
)

func (c *ControllerV1) Denglu(ctx context.Context, req *v1.DengluReq) (res *v1.DengluRes, err error) {
	err = c.users.Denglu(ctx, req.Username, req.Password)
	return nil, err
}
