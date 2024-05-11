package account_payment

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"be-assesment/src/app"
	"be-assesment/src/modules/payment_manager_services/entity"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
func TestAccountPayment_UpdateAccountPayment(t *testing.T) {
	query := `UPDATE account_payments SET amount = amount + $1, updated_at = $2 WHERE id = $3 RETURNING id, name, type, user_id, amount`

	tests := []struct {
		name    string
		ap      *AccountPayment
		want    entity.AccountPayment
		wantErr bool
		setup   func(mock sqlmock.Sqlmock)
	}{
		{
			name:    "Success",
			want:    entity.AccountPayment{},
			wantErr: false,
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				rows := sqlmock.NewRows([]string{"id", "name", "type", "user_id", "amount"}).
					AddRow("w", "halo", "makassar", "k", float64(9))
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(float64(10), AnyTime{}, "s").WillReturnRows(rows)
			},
		},
		{
			name:    "Failure - DB Error",
			wantErr: true,
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs(float64(10), AnyTime{}, "s").
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

			td := &AccountPayment{
				app.Db{Master: db, Slave: db},
			}
			// Set up the mock expectations
			tt.setup(mocks)

			tx, _ := db.Begin()

			// Call the function
			_, err = td.UpdateAccountPayment(context.Background(), tx, entity.AccountPayment{Amount: float64(10), Id: "s"})

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
