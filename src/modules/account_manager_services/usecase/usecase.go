package usecase

import (
	"context"

	"be-assesment/src/app"
	"be-assesment/src/modules/account_manager_services/entity"
	"be-assesment/src/modules/account_manager_services/repository"
)

type (
	UseCaseInterface interface {
		SignUp(context.Context, entity.SingUpReq) (string, error)
		SignIn(ctx context.Context, req entity.User) (string, error)
		GetAllUsersWithPayments(ctx context.Context) ([]entity.UserWithPayments, error)
	}

	UseCase struct {
		app  *app.App
		repo *repository.Repository
	}
)

func Init(app *app.App, repo *repository.Repository) UseCaseInterface {
	return &UseCase{
		app:  app,
		repo: repo,
	}
}
