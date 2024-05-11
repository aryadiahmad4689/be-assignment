package repository

import (
	"be-assesment/src/app"
	"be-assesment/src/modules/payment_manager_services/repository/account_payment"
	"be-assesment/src/modules/payment_manager_services/repository/transaction"
	"context"
	"database/sql"
)

type (
	Repository struct {
		app            *app.App
		Transaction    transaction.TransactionInterface
		AccountPayment account_payment.AccountPaymentInterface
	}
)

func Init(app *app.App) *Repository {
	return &Repository{
		Transaction:    transaction.Init(app),
		AccountPayment: account_payment.Init(app),
		app:            app,
	}
}

func (r *Repository) BeginTx(ctx context.Context) (*sql.Tx, error) {
	var db = r.app.GetDb()
	var transaction, err = db.Master.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}
