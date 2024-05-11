package transaction

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"be-assesment/src/modules/payment_manager_services/entity"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func nullStringFrom(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

func TestTransaction_LogTransaction(t *testing.T) {
	ctx := context.Background()
	transaction := entity.HistoricalTransaction{
		AccountPaymentIdPayed:    nullStringFrom("eae581cd-7f29-4d12-bd70-798466de36a0"),
		AccountPaymentIdReceived: nullStringFrom("eae581cd-7f29-4d12-bd70-798466de36a0"),
		TotalTransaction:         100.0,
		Currency:                 "USD",
		StatusTransaction:        "completed",
		TypeTransaction:          "transfer",
	}
	query := `INSERT INTO historical_transactions (account_payment_id_payed,account_payment_id_received, total_transaction, currency, status_transaction, type_transaction, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, account_payment_id_payed,account_payment_id_received, total_transaction, currency, status_transaction, type_transaction`

	tests := []struct {
		name    string
		trx     *Transaction
		wantErr bool
		setup   func(mock sqlmock.Sqlmock)
	}{
		{
			name: "Success",

			wantErr: false,
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				rows := sqlmock.NewRows([]string{"id", "account_payment_id_payed", "account_payment_id_received", "total_transaction", "currency", "status_transaction", "type_transaction"}).
					AddRow(1, "1", "2", 100.0, "USD", "completed", "transfer")
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(transaction.AccountPaymentIdPayed, transaction.AccountPaymentIdReceived, transaction.TotalTransaction, transaction.Currency, transaction.StatusTransaction, transaction.TypeTransaction, AnyTime{}, AnyTime{}).WillReturnRows(rows)
			},
		},
		{
			name: "Failure - DB Error",

			wantErr: true,
			setup: func(mocks sqlmock.Sqlmock) {

				mocks.ExpectBegin()
				mocks.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(transaction.AccountPaymentIdPayed, transaction.AccountPaymentIdReceived, transaction.TotalTransaction, transaction.Currency, transaction.StatusTransaction, transaction.TypeTransaction, AnyTime{}, AnyTime{}).
					WillReturnError(sql.ErrNoRows) // Simulate a connection error
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mocks, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			trx := &Transaction{}

			// Set up the mock expectations
			tt.setup(mocks)
			tx, _ := db.Begin()

			// Call the function
			_, err = trx.LogTransaction(ctx, tx, transaction)

			// Assertions
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mocks.ExpectationsWereMet())
		})
	}
}
