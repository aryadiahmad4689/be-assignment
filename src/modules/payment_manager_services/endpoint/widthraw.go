package endpoint

import (
	"be-assesment/src/modules/payment_manager_services/entity"
	"context"

	"github.com/go-playground/validator/v10"
)

func (ed *Endpoint) Widthdraw(ctx context.Context, wdthReq WithdrawReq) (interface{}, error) {
	validate := validator.New()

	err := validate.Struct(wdthReq)
	if err != nil {
		return SendResp{}, err
	}
	_, err = ed.usecase.Withdraw(ctx, entity.WithdrawReq{
		AccountPaymentId: wdthReq.AccountPaymentPayed,
		TotalTransaction: wdthReq.TotalTransaction,
		Currency:         "USD",  // for now just USD
		Type:             "WDTH", // harcode for now
	})
	if err != nil {
		return SendResp{}, err
	}
	return SendResp{}, nil
}
