package usecase

import (
	"context"
	"errors"
	"os"
	"testing"

	"be-assesment/src/app"
	"be-assesment/src/modules/account_manager_services/entity"
	"be-assesment/src/modules/account_manager_services/repository"
	mockAcc "be-assesment/src/modules/account_manager_services/repository/account_payment/mocks"
	"be-assesment/src/modules/account_manager_services/repository/user/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestUseCase_SignUp(t *testing.T) {
	var apps = &app.App{}
	db, _, _ := sqlmock.New()
	os.Setenv("AUTH_KEY", "JWTSecret") // Set test secret key

	type args struct {
		ctx  context.Context
		user entity.SingUpReq
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockDb  func(a args) *repository.Repository
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				user: entity.SingUpReq{
					Name:     "Test User",
					Email:    "test@example.com",
					Password: "password123",
				},
			},
			wantErr: false,
			mockDb: func(a args) *repository.Repository {
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(a.user.Password), bcrypt.DefaultCost)
				dbs := &app.Db{
					Master: db,
					Slave:  db,
				}
				apps.SetDb(dbs)
				mockUser := new(mocks.UserInterfaceMock)
				mockAccountPayment := new(mockAcc.AccountUserInterfaceMock)
				mockUser.On("StoreUser", mock.Anything, mock.Anything).Return(entity.User{
					Id:       "123",
					Password: string(hashedPassword),
				}, nil)
				mockAccountPayment.On("StoreAccountPayment", mock.Anything, mock.Anything).Return(entity.AccountPayment{}, nil)
				repo := repository.Init(apps).SetUser(mockUser).SetAccountPayment(mockAccountPayment)
				return repo
			},
		},
		{
			name: "fail store account payment",
			args: args{
				ctx: context.Background(),
				user: entity.SingUpReq{
					Name:     "Test User",
					Email:    "test@example.com",
					Password: "password123",
				},
			},
			wantErr: true,
			mockDb: func(a args) *repository.Repository {
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(a.user.Password), bcrypt.DefaultCost)
				dbs := &app.Db{
					Master: db,
					Slave:  db,
				}
				apps.SetDb(dbs)
				mockUser := new(mocks.UserInterfaceMock)
				mockAccountPayment := new(mockAcc.AccountUserInterfaceMock)
				mockUser.On("StoreUser", mock.Anything, mock.Anything).Return(entity.User{
					Id:       "123",
					Password: string(hashedPassword),
				}, nil)
				mockAccountPayment.On("StoreAccountPayment", mock.Anything, mock.Anything).Return(entity.AccountPayment{}, errors.New("err"))
				repo := repository.Init(apps).SetUser(mockUser).SetAccountPayment(mockAccountPayment)
				return repo
			},
		},
		{
			name: "fail - error hashing password",
			args: args{
				ctx: context.Background(),
				user: entity.SingUpReq{
					Name:  "Test User",
					Email: "test@example.com",
				},
			},
			wantErr: true,
			mockDb: func(a args) *repository.Repository {
				// Mock setup
				// ...
				dbs := &app.Db{
					Master: db,
					Slave:  db,
				}
				apps.SetDb(dbs)
				mockUser := new(mocks.UserInterfaceMock)
				mockUser.On("StoreUser", mock.Anything, mock.Anything).Return(entity.User{}, errors.New("hashing error"))
				repo := repository.Init(apps).SetUser(mockUser)
				return repo
			},
		},
		{
			name: "fail - error storing user",
			args: args{
				ctx: context.Background(),
				user: entity.SingUpReq{
					Name:     "Test User",
					Email:    "test@example.com",
					Password: "password123",
				},
			},
			wantErr: true,
			mockDb: func(a args) *repository.Repository {
				// Mock setup
				// ...
				dbs := &app.Db{
					Master: db,
					Slave:  db,
				}
				apps.SetDb(dbs)
				mockUser := new(mocks.UserInterfaceMock)
				mockUser.On("StoreUser", mock.Anything, mock.AnythingOfType("entity.User")).Return(entity.User{}, errors.New("database error"))
				repo := repository.Init(apps).SetUser(mockUser)
				return repo
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.mockDb(tt.args)
			u := Init(apps, repo)
			token, err := u.SignUp(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.SignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				assert.NotEmpty(t, token)
			}
		})
	}
}
