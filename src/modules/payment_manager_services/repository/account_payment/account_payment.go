package account_payment

import (
	"context"
	"database/sql"

	"be-assesment/src/app"
	"be-assesment/src/modules/payment_manager_services/entity"
)

type AccountPayment struct {
	db app.Db
}

type AccountPaymentInterface interface {
	UpdateAccountPayment(ctx context.Context, trx *sql.Tx, accountPayment entity.AccountPayment) (entity.AccountPayment, error)
}

func Init(app *app.App) AccountPaymentInterface {
	var db = app.GetDb()
	return &AccountPayment{
		db: *db,
	}
}
