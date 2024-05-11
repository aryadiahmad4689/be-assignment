package endpoint

import (
	"be-assesment/src/app"
	"be-assesment/src/modules/payment_manager_services/usecase"
	"context"
)

type (
	WithdrawReq struct {
		AccountPaymentPayed string  `json:"account_payment_id_withdraw" validate:"required"`
		TotalTransaction    float64 `json:"total_transaction" validate:"required"`
		Currency            string
		Type                string
	}
	WithdrawResp struct{}

	SendReq struct {
		AccountPaymentIdPayed    string  `json:"account_payment_id_payed" validate:"required" `
		AccountPaymentIdReceived string  `json:"account_payment_id_received" validate:"required"`
		TotalTransaction         float64 `json:"total_transaction" validate:"required"`
		Type                     string
		Currency                 string
	}

	SendResp struct{}

	EndpointInterface interface {
		Widthdraw(context.Context, WithdrawReq) (interface{}, error)
		Send(context.Context, SendReq) (interface{}, error)
	}
	Endpoint struct {
		app     *app.App
		usecase usecase.UseCaseInterface
	}
)

func Init(app *app.App, usecase usecase.UseCaseInterface) EndpointInterface {
	return &Endpoint{
		app:     app,
		usecase: usecase,
	}
}
