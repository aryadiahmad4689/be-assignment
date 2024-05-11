package usecase

import (
	"context"
	"os"
	"time"

	"be-assesment/src/modules/account_manager_services/entity"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func (us *UseCase) SignUp(ctx context.Context, singUpReq entity.SingUpReq) (string, error) {
	// Hash password
	user := entity.User{
		Name:       singUpReq.Name,
		Age:        singUpReq.Age,
		Gender:     singUpReq.Gender,
		Email:      singUpReq.Email,
		Password:   singUpReq.Password,
		IsVerified: singUpReq.IsVerified, // I consider all those listed as verified for now
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashedPassword)

	newUser, err := us.repo.User.StoreUser(ctx, user)
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(10 * time.Minute)

	// claim token
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   newUser.Id,
	}

	// create new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sing token
	tokenString, err := token.SignedString([]byte(os.Getenv("AUTH_KEY")))
	if err != nil {
		return "", err
	}

	// insert account payment
	accountPay := entity.AccountPayment{
		Name:   singUpReq.NameAccount,
		Type:   singUpReq.Type,
		UserId: newUser.Id,
		Amount: float64(100000), //I assume everyone who registers has 1000 dollars in their account
	}
	_, err = us.repo.AccountPayment.StoreAccountPayment(ctx, accountPay)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
