package account_payment

import (
	"context"
	"database/sql"
	"time"

	"be-assesment/src/modules/payment_manager_services/entity"
)

func (u *AccountPayment) UpdateAccountPayment(ctx context.Context, trx *sql.Tx, accountPayment entity.AccountPayment) (entity.AccountPayment, error) {
	// The SQL query is modified to update the 'amount' and 'updated_at' fields where the 'id' matches.
	query := `UPDATE account_payments SET amount = amount + $1, updated_at = $2 WHERE id = $3 RETURNING id, name, type, user_id, amount`

	var now = time.Now().UTC()

	// Execute the update statement and return the updated row.
	err := trx.QueryRowContext(ctx, query, accountPayment.Amount, now, accountPayment.Id).
		Scan(&accountPayment.Id, &accountPayment.Name, &accountPayment.Type, &accountPayment.UserId, &accountPayment.Amount)
	if err != nil {
		return entity.AccountPayment{}, err
	}

	return accountPayment, nil
}
