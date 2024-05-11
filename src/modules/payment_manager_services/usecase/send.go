package usecase

import (
	"be-assesment/src/modules/payment_manager_services/entity"
	"context"
)

func (us *UseCase) Send(ctx context.Context, transaction entity.SendReq) (entity.SendReq, error) {

	tx, err := us.repo.BeginTx(ctx)
	if err != nil {
		return entity.SendReq{}, err
	}
	// Decrement amount from the payer's account
	accountPayed := entity.AccountPayment{
		Id:     transaction.AccountPaymentIdPayed,
		Amount: -transaction.TotalTransaction,
	}
	if _, err = us.repo.AccountPayment.UpdateAccountPayment(ctx, tx, accountPayed); err != nil {
		return entity.SendReq{}, err
	}

	// Increment amount to the receiver's account
	accountReceived := entity.AccountPayment{
		Id:     transaction.AccountPaymentIdReceived,
		Amount: transaction.TotalTransaction,
	}
	if _, err = us.repo.AccountPayment.UpdateAccountPayment(ctx, tx, accountReceived); err != nil {
		return entity.SendReq{}, err
	}

	historicalTransaction := entity.HistoricalTransaction{
		AccountPaymentIdPayed:    nullStringFrom(transaction.AccountPaymentIdPayed),
		AccountPaymentIdReceived: nullStringFrom(transaction.AccountPaymentIdReceived),
		TotalTransaction:         transaction.TotalTransaction,
		TypeTransaction:          transaction.Type,
		Currency:                 transaction.Currency, //for now usd
	}

	if _, err = us.repo.Transaction.LogTransaction(ctx, tx, historicalTransaction); err != nil {
		return entity.SendReq{}, err
	}

	tx.Commit()
	return transaction, nil
}
