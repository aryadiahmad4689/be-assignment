package usecase

import (
	"be-assesment/src/app"
	"be-assesment/src/modules/payment_manager_services/entity"
	"be-assesment/src/modules/payment_manager_services/repository"
	"context"
)

type (
	UseCaseInterface interface {
		Withdraw(ctx context.Context, transaction entity.WithdrawReq) (entity.WithdrawReq, error)
		Send(ctx context.Context, transaction entity.SendReq) (entity.SendReq, error)
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
