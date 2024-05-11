package mocks

import (
	"context"

	"be-assesment/src/modules/account_manager_services/entity"
	"be-assesment/src/modules/account_manager_services/repository/account_payment"

	"github.com/stretchr/testify/mock"
)

type AccountUserInterfaceMock struct {
	mock.Mock
}

func Init() account_payment.AccountPaymentInterface {
	return &AccountUserInterfaceMock{}
}

func (m *AccountUserInterfaceMock) StoreAccountPayment(ctx context.Context, accountPayment entity.AccountPayment) (entity.AccountPayment, error) {
	ret := m.Called(ctx, accountPayment)
	return ret.Get(0).(entity.AccountPayment), ret.Error(1)
}
