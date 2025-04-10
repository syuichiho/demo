// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UsersDao is the data access object for the table users.
type UsersDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of the current DAO.
	columns UsersColumns // columns contains all the column names of Table for convenient usage.
}

// UsersColumns defines and stores column names for the table users.
type UsersColumns struct {
	Id           string // 主键，自动递增，唯一标识用户
	Username     string // 用户名
	Password     string // 密码
	Email        string // 邮箱
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
	RealName     string // 真实姓名
	Phone        string // 电话号码
	Avatar       string // 头像
	DepartmentId string // 部门id
	RoleId       string // 角色id
	Status       string //
}

// usersColumns holds the columns for the table users.
var usersColumns = UsersColumns{
	Id:           "id",
	Username:     "username",
	Password:     "password",
	Email:        "email",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	RealName:     "real_name",
	Phone:        "phone",
	Avatar:       "avatar",
	DepartmentId: "department_id",
	RoleId:       "role_id",
	Status:       "status",
}

// NewUsersDao creates and returns a new DAO object for table data access.
func NewUsersDao() *UsersDao {
	return &UsersDao{
		group:   "default",
		table:   "users",
		columns: usersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UsersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UsersDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UsersDao) Columns() UsersColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UsersDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UsersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UsersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
