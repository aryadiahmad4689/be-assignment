package usecase

import (
	"be-assesment/src/modules/payment_manager_services/entity"
	"context"
	"fmt"
)

func (us *UseCase) Withdraw(ctx context.Context, transaction entity.WithdrawReq) (entity.WithdrawReq, error) {
	tx, err := us.repo.BeginTx(ctx)
	if err != nil {
		return entity.WithdrawReq{}, err
	}
	// Decrement amount from the account
	accountToWithdraw := entity.AccountPayment{
		Id:     transaction.AccountPaymentId,
		Amount: -transaction.TotalTransaction, // Negative because we are reducing the balance
	}

	if _, err := us.repo.AccountPayment.UpdateAccountPayment(ctx, tx, accountToWithdraw); err != nil {
		return entity.WithdrawReq{}, err
	}

	// Log the transaction
	historicalTransaction := entity.HistoricalTransaction{
		AccountPaymentIdPayed: nullStringFrom(transaction.AccountPaymentId), // Assuming this is the account from which funds are withdrawn
		TotalTransaction:      transaction.TotalTransaction,
		TypeTransaction:       transaction.Type, // widthdraw
		Currency:              transaction.Currency,
	}

	if _, err := us.repo.Transaction.LogTransaction(ctx, tx, historicalTransaction); err != nil {
		fmt.Println("a", err)
		return entity.WithdrawReq{}, err
	}

	return transaction, nil
}
