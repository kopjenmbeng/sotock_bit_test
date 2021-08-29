package products

import (
	"net/http"

	"github.com/kopjenmbeng/sotock_bit_test/internal/utility/validator"
)

type AddProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Qty   int     `json:"qty"`
}

type UpdateProductRequest struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Qty   int     `json:"qty"`
}

func (req *AddProductRequest) Bind(r *http.Request) error {
	if err := req.Validate(r); err != nil {
		return err
	}
	return nil
}

func (req *AddProductRequest) Validate(r *http.Request) error {
	if err := validator.ValidateEmpty("name", req.Name); err != nil {
		return err
	}
	if err := validator.ValidateEmpty("price", req.Price); err != nil {
		return err
	}
	if err := validator.ValidateEmpty("qty", req.Qty); err != nil {
		return err
	}

	return nil
}

func (req *UpdateProductRequest) Bind(r *http.Request) error {
	if err := req.Validate(r); err != nil {
		return err
	}
	return nil
}

func (req *UpdateProductRequest) Validate(r *http.Request) error {
	if err := validator.ValidateEmpty("id", req.Id); err != nil {
		return err
	}
	if err := validator.ValidateEmpty("name", req.Name); err != nil {
		return err
	}
	if err := validator.ValidateEmpty("price", req.Price); err != nil {
		return err
	}
	if err := validator.ValidateEmpty("qty", req.Qty); err != nil {
		return err
	}

	return nil
}
