package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ZhuceReq struct {
	g.Meta   `path:"/zhuce" method:"post"`
	Username string `json:"username" v:"required#用户名不能为空"`
	Password string `json:"password" v:"required#密码不能为空"`
}

type ZhuceRes struct {
}
