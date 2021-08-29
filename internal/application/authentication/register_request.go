package authentication

import (
	"net/http"

	"github.com/kopjenmbeng/sotock_bit_test/internal/utility/validator"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}

func (req *RegisterRequest) Bind(r *http.Request) error {
	if err := req.Validate(r); err != nil {
		return err
	}
	return nil
}

func (req *RegisterRequest) Validate(r *http.Request) error {
	if err := validator.ValidateEmpty("email", req.Email); err != nil {
		return err
	}
	if err := validator.ValidateEmpty("full_name", req.FullName); err != nil {
		return err
	}
	if err := validator.ValidateEmpty("password", req.Password); err != nil {
		return err
	}

	return nil
}
