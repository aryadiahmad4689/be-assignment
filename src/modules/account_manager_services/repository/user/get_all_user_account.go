package user

import (
	"be-assesment/src/modules/account_manager_services/entity"
	"context"
	"log"
)

func (us *User) GetAllUsersWithPayments(ctx context.Context) ([]entity.UserWithPayments, error) {
	query := `
SELECT 
    u.id, u.name, u.email,
    ap.id AS payment_id, ap.name AS payment_name, ap.type, ap.amount
FROM users u
LEFT JOIN account_payments ap ON u.id = ap.user_id;
`

	rows, err := us.db.Slave.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users map[string]*entity.UserWithPayments = make(map[string]*entity.UserWithPayments)

	for rows.Next() {
		var user entity.UserWithPayments
		var payment entity.AccountPayment
		err = rows.Scan(
			&user.Id, &user.Name, &user.Email,
			&payment.Id, &payment.Name, &payment.Type, &payment.Amount)
		if err != nil {
			log.Println("Failed to scan row:", err)
			continue
		}

		if usr, ok := users[user.Id]; ok {
			// user already exists, just append the payment
			if payment.Id != "" {
				usr.Payments = append(usr.Payments, payment)
			}
		} else {
			// new user, add to map
			if payment.Id != "" {
				user.Payments = append(user.Payments, payment)
			}
			users[user.Id] = &user
		}
	}

	// Convert map to slice
	var userList []entity.UserWithPayments
	for _, user := range users {
		userList = append(userList, *user)
	}

	return userList, nil
}
