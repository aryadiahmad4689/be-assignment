package modules

import (
	"be-assesment/src/app"
	accManagerModule "be-assesment/src/modules/account_manager_services"
	payManagerModule "be-assesment/src/modules/payment_manager_services"

	"be-assesment/src/modules/interfaces"
)

type Modules struct {
	AccountManager interfaces.ModuleInterface
	PaymentManager interfaces.ModuleInterface
}

func Init(app *app.App) *Modules {
	return &Modules{
		AccountManager: accManagerModule.Init(app),
		PaymentManager: payManagerModule.Init(app),
	}
}
