package users

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"demo/api/users/v1"
)

func (c *ControllerV1) Zhuce(ctx context.Context, req *v1.ZhuceReq) (res *v1.ZhuceRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
