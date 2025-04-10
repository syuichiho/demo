// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure of table users for DAO operations like Where/Data.
type Users struct {
	g.Meta       `orm:"table:users, do:true"`
	Id           interface{} // 主键，自动递增，唯一标识用户
	Username     interface{} // 用户名
	Password     interface{} // 密码
	Email        interface{} // 邮箱
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	RealName     interface{} // 真实姓名
	Phone        interface{} // 电话号码
	Avatar       interface{} // 头像
	DepartmentId interface{} // 部门id
	RoleId       interface{} // 角色id
	Status       interface{} //
}
