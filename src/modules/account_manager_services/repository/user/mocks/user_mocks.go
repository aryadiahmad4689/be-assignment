package mocks

import (
	"context"

	"be-assesment/src/modules/account_manager_services/entity"
	"be-assesment/src/modules/account_manager_services/repository/user"

	"github.com/stretchr/testify/mock"
)

type UserInterfaceMock struct {
	mock.Mock
}

func Init() user.UserInterface {
	return &UserInterfaceMock{}
}

func (m *UserInterfaceMock) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	ret := m.Called(ctx, email)
	return ret.Get(0).(entity.User), ret.Error(1)
}

func (m *UserInterfaceMock) StoreUser(ctx context.Context, user entity.User) (entity.User, error) {
	ret := m.Called(ctx, user)
	return ret.Get(0).(entity.User), ret.Error(1)
}

func (m *UserInterfaceMock) GetAllUsersWithPayments(ctx context.Context) ([]entity.UserWithPayments, error) {
	ret := m.Called(ctx)
	return ret.Get(0).([]entity.UserWithPayments), ret.Error(1)
}

