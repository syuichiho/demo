package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DengluReq struct {
	g.Meta   `path:"/users" method:"post"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type DengluRes struct {
}
