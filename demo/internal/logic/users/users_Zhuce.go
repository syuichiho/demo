package users

import (
	"context"
	"demo/internal/dao"
	"demo/internal/model/do"
)

func (u *Users) Zhuce(ctx context.Context, username, password string) error {
	_, err := dao.Users.Ctx(ctx).Data(do.Users{
		Username: username,
		Password: password,
	}).Insert()

	if err != nil {
		return err
	}
	return nil

}
