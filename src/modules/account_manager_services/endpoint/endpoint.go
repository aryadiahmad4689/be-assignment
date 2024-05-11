package endpoint

import (
	"context"

	"be-assesment/src/app"
	"be-assesment/src/modules/account_manager_services/usecase"
)

type (
	EndpointInterface interface {
		SignUp(context.Context, SignUpRequest) (interface{}, error)
		SignIn(ctx context.Context, req SignInRequest) (interface{}, error)
		GetAllUsersWithPayments(ctx context.Context) (interface{}, error)
	}
	Endpoint struct {
		app     *app.App
		usecase usecase.UseCaseInterface
	}

	SignUpRequest struct {
		Name        string `json:"name" validate:"required"`
		Email       string `json:"email" validate:"required"`
		Password    string `json:"password" validate:"required"`
		Age         int    `json:"age" validate:"gte=0,lte=130"`
		Gender      string `json:"gender" validate:"required"`
		NameAccount string `json:"account_name" validate:"required"`
		Type        string `json:"type" validate:"required"`
	}

	SignUpResp struct {
		Token string `json:"token"`
	}

	SignInRequest struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	SignInResp struct {
		Token string `json:"token"`
	}

	RepNull struct{}
)

func Init(app *app.App, usecase usecase.UseCaseInterface) EndpointInterface {
	return &Endpoint{
		app:     app,
		usecase: usecase,
	}
}
