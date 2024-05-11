package user

import (
	"context"

	"be-assesment/src/app"
	"be-assesment/src/modules/account_manager_services/entity"
)

type User struct {
	db app.Db
}

type UserInterface interface {
	StoreUser(ctx context.Context, user entity.User) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	GetAllUsersWithPayments(ctx context.Context) ([]entity.UserWithPayments, error)
}

func Init(app *app.App) UserInterface {
	var db = app.GetDb()
	return &User{
		db: *db,
	}
}
