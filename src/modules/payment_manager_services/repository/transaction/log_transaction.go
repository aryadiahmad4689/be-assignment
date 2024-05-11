package transaction

import (
	"be-assesment/src/modules/payment_manager_services/entity"
	"context"
	"database/sql"
	"time"
)

func (u *Transaction) LogTransaction(ctx context.Context, trx *sql.Tx, transaction entity.HistoricalTransaction) (entity.HistoricalTransaction, error) {
	query := `INSERT INTO historical_transactions (account_payment_id_payed,account_payment_id_received, total_transaction, currency, status_transaction, type_transaction, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, account_payment_id_payed,account_payment_id_received, total_transaction, currency, status_transaction, type_transaction`

	var now = time.Now().UTC()
	var data = entity.HistoricalTransaction{}

	err := trx.QueryRowContext(ctx, query,
		transaction.AccountPaymentIdPayed, transaction.AccountPaymentIdReceived, transaction.TotalTransaction, transaction.Currency, transaction.StatusTransaction, transaction.TypeTransaction, now, now).
		Scan(&data.Id, &data.AccountPaymentIdPayed, &data.AccountPaymentIdReceived, &data.TotalTransaction, &data.Currency, &data.StatusTransaction, &data.TypeTransaction)
	if err != nil {
		return entity.HistoricalTransaction{}, err
	}

	return data, nil
}
