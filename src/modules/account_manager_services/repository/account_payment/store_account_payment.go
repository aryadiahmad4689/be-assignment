package account_payment

import (
	"context"
	"time"

	"be-assesment/src/modules/account_manager_services/entity"
)

func (u *AccountPayment) StoreAccountPayment(ctx context.Context, accountPayment entity.AccountPayment) (entity.AccountPayment, error) {
	query := `INSERT INTO account_payments (name, type, user_id, amount,created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, type, user_id, amount`

	var now = time.Now().UTC()
	var data = entity.AccountPayment{}

	err := u.db.Master.QueryRowContext(ctx, query,
		accountPayment.Name, accountPayment.Type, accountPayment.UserId, accountPayment.Amount, now, now).
		Scan(&data.Id, &data.Name, &data.Type, &data.UserId, &data.Amount)
	if err != nil {
		return entity.AccountPayment{}, err
	}

	return data, nil
}
