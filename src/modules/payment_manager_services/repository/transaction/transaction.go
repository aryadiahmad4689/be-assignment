package transaction

import (
	"context"
	"database/sql"

	"be-assesment/src/app"
	"be-assesment/src/modules/payment_manager_services/entity"
)

type Transaction struct {
	db app.Db
}

type TransactionInterface interface {
	LogTransaction(ctx context.Context, trx *sql.Tx, user entity.HistoricalTransaction) (entity.HistoricalTransaction, error)
}

func Init(app *app.App) TransactionInterface {
	var db = app.GetDb()
	return &Transaction{
		db: *db,
	}
}
