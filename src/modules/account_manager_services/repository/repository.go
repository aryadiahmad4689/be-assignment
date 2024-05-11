package repository

import (
	"be-assesment/src/app"
	"be-assesment/src/modules/account_manager_services/repository/account_payment"
	"be-assesment/src/modules/account_manager_services/repository/user"
)

type (
	Repository struct {
		app            *app.App
		User           user.UserInterface
		AccountPayment account_payment.AccountPaymentInterface
	}
)

func Init(app *app.App) *Repository {
	return &Repository{
		app:            app,
		User:           user.Init(app),
		AccountPayment: account_payment.Init(app),
	}
}

func (r *Repository) SetUser(u user.UserInterface) *Repository {
	r.User = u
	return r
}
func (r *Repository) SetAccountPayment(u account_payment.AccountPaymentInterface) *Repository {
	r.AccountPayment = u
	return r
}
