package endpoint

import (
	"be-assesment/src/modules/payment_manager_services/entity"
	"context"

	"github.com/go-playground/validator/v10"
)

func (ed *Endpoint) Send(ctx context.Context, sendReq SendReq) (interface{}, error) {
	validate := validator.New()

	err := validate.Struct(sendReq)
	if err != nil {
		return SendResp{}, err
	}
	_, err = ed.usecase.Send(ctx, entity.SendReq{
		AccountPaymentIdPayed:    sendReq.AccountPaymentIdPayed,
		AccountPaymentIdReceived: sendReq.AccountPaymentIdReceived,
		TotalTransaction:         sendReq.TotalTransaction,
		Currency:                 "USD",  // for now just USD
		Type:                     "Send", // harcode for now
	})
	if err != nil {
		return SendResp{}, err
	}
	return SendResp{}, nil
}
