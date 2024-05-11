package account_payment

import (
	"context"

	"be-assesment/src/app"
	"be-assesment/src/modules/account_manager_services/entity"
)

type AccountPayment struct {
	db app.Db
}

type AccountPaymentInterface interface {
	StoreAccountPayment(ctx context.Context, accountPayment entity.AccountPayment) (entity.AccountPayment, error)
}

func Init(app *app.App) AccountPaymentInterface {
	var db = app.GetDb()
	return &AccountPayment{
		db: *db,
	}
}
