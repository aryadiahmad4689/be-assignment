package endpoint

import (
	"context"

	"be-assesment/src/modules/account_manager_services/entity"

	"github.com/go-playground/validator/v10"
)

func (ed *Endpoint) SignIn(ctx context.Context, req SignInRequest) (interface{}, error) {
	validate := validator.New()

	err := validate.Struct(req)
	if err != nil {
		return SignInResp{}, err
	}
	
	token, err := ed.usecase.SignIn(ctx, entity.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return SignInResp{}, err
	}
	return SignInResp{
		Token: token,
	}, nil
}
