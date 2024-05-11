package usecase

import (
	"be-assesment/src/modules/account_manager_services/entity"
	"context"
)

func (us *UseCase) GetAllUsersWithPayments(ctx context.Context) ([]entity.UserWithPayments, error) {
	data, err := us.repo.User.GetAllUsersWithPayments(ctx)
	if err != nil {
		return []entity.UserWithPayments{}, err
	}

	return data, nil
}
