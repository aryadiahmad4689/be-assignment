package user

import (
	"context"

	"be-assesment/src/modules/account_manager_services/entity"
)

func (u *User) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	query := `SELECT id, email, password, name, age, gender 
	          FROM users WHERE email = $1`

	err := u.db.Slave.QueryRowContext(ctx, query, email).Scan(
		&user.Id, &user.Email, &user.Password, &user.Name, &user.Age, &user.Gender)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
