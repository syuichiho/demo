// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	Id           uint        `json:"id"           orm:"id"            description:"主键，自动递增，唯一标识用户"` // 主键，自动递增，唯一标识用户
	Username     string      `json:"username"     orm:"username"      description:"用户名"`            // 用户名
	Password     string      `json:"password"     orm:"password"      description:"密码"`             // 密码
	Email        string      `json:"email"        orm:"email"         description:"邮箱"`             // 邮箱
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`           // 创建时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`           // 更新时间
	RealName     string      `json:"realName"     orm:"real_name"     description:"真实姓名"`           // 真实姓名
	Phone        string      `json:"phone"        orm:"phone"         description:"电话号码"`           // 电话号码
	Avatar       string      `json:"avatar"       orm:"avatar"        description:"头像"`             // 头像
	DepartmentId uint        `json:"departmentId" orm:"department_id" description:"部门id"`           // 部门id
	RoleId       uint        `json:"roleId"       orm:"role_id"       description:"角色id"`           // 角色id
	Status       int         `json:"status"       orm:"status"        description:""`               //
}
