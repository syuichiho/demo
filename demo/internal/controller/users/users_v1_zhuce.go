package users

import (
	"context"

	v1 "demo/api/users/v1"
)

func (c *ControllerV1) Zhuce(ctx context.Context, req *v1.ZhuceReq) (res *v1.ZhuceRes, err error) {
	err = c.users.Zhuce(ctx,
		req.Username,
		req.Password,
	)
	return nil, err
}
