package authentication

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/kopjenmbeng/sotock_bit_test/internal/dto"
	"github.com/kopjenmbeng/sotock_bit_test/internal/middleware"
	"github.com/kopjenmbeng/sotock_bit_test/internal/middleware/jwe_auth"
)

type IAuthenticationUseCase interface {
	GetToken(ctx context.Context, email string, password string) (token *TokenResponse, status int, Err error)
	Register(ctx context.Context, req RegisterRequest) (status int, err error)
}

type AuthenticationUseCase struct {
	repository IAuthenticationRepository
	r          *http.Request
}

func NewAuthenticationUseCase(repository IAuthenticationRepository, r *http.Request) IAuthenticationUseCase {
	return &AuthenticationUseCase{repository: repository, r: r}
}

func (use_case *AuthenticationUseCase) Register(ctx context.Context, req RegisterRequest) (status int, err error) {

	salt := uuid.New().String()
	iteration := middleware.GenerateRandomNumber(900, 999)
	secLength := middleware.GenerateRandomNumber(32, 64)
	password := middleware.HashPassword(req.Password, salt, iteration, secLength)
	cus := dto.User{Id: uuid.New().String(), Email: req.Email, FullName: req.FullName, Salt: salt, Password: password, Iteration: iteration, SecurityLength: secLength}
	status, err = use_case.repository.Register(ctx, cus)
	return
}
func (use_case *AuthenticationUseCase) GetToken(ctx context.Context, email string, password string) (token *TokenResponse, status int, Err error) {
	data, status, err := use_case.repository.GetUser(ctx, email)
	if err != nil {
		return nil, status, err
	}
	var result TokenResponse
	if status == http.StatusOK {
		paramPassword := middleware.HashPassword(password, data.Salt, data.Iteration, data.SecurityLength)
		if data.Password != paramPassword {
			return nil, http.StatusUnauthorized, errors.New("password yang anda masukan salah.")
		}
		token, Exp, _ := jwe_auth.GenerateToken(use_case.r, data.Id, false)
		result.Token = token
		result.Expiry = Exp

	}
	return &result, 0, nil
}
