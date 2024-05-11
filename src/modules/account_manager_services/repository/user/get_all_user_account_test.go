package user

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"be-assesment/src/modules/account_manager_services/entity"

	"be-assesment/src/app"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestUser_GetAllUsersWithPayments(t *testing.T) {
	query := `
	SELECT 
		u.id, u.name, u.email,
		ap.id AS payment_id, ap.name AS payment_name, ap.type, ap.amount
	FROM users u
	LEFT JOIN account_payments ap ON u.id = ap.user_id;
	`
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		us      *User
		args    args
		want    []entity.UserWithPayments
		wantErr bool
		setup   func(mock sqlmock.Sqlmock)
	}{
		{
			name: "Success",
			args: args{
				ctx: context.TODO(),
			},
			want: []entity.UserWithPayments{
				{},
			},
			wantErr: false,
			setup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "amount", "date"}).
					AddRow(1, "John Doe", 100.0, "2021-01-01")
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WillReturnRows(rows)
			},
		},
		{
			name: "Failure - DB Error",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).
					WillReturnError(sql.ErrConnDone) // Simulate a connection error
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			td := &User{
				db: app.Db{Slave: db},
			}

			// Set up the mock expectations
			tt.setup(mock)

			// Call the function
			_, err = td.GetAllUsersWithPayments(tt.args.ctx)

			// Assertions
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
