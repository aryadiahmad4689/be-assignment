package endpoint

import (
	"context"

	"be-assesment/src/modules/account_manager_services/entity"

	"github.com/go-playground/validator/v10"
)

// i conside when sign up inser all field for now
func (ed *Endpoint) SignUp(ctx context.Context, req SignUpRequest) (interface{}, error) {

	validate := validator.New()

	err := validate.Struct(req)
	if err != nil {
		return SignInResp{}, err
	}
	token, err := ed.usecase.SignUp(ctx, entity.SingUpReq{
		Name:        req.Name,
		Email:       req.Email,
		Password:    req.Password,
		Age:         req.Age,
		Gender:      req.Gender,
		IsVerified:  0, // I consider all those listed as verified for now , 0 its mean verified
		NameAccount: req.NameAccount,
		Type:        req.Type,
	})
	if err != nil {
		return SignUpResp{}, err
	}
	return SignUpResp{
		Token: token,
	}, nil
}
