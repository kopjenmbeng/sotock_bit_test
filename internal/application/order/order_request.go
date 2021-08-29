package order

import (
	"net/http"

	"github.com/kopjenmbeng/sotock_bit_test/internal/utility/validator"
)

type CreateOrderRequest struct {
	ProductId string `json:"product_id"`
	Qty       int    `json:"qty"`
}

type PaidOrderRequest struct {
	OrderId string `json:"order_id"`
}

func (req *CreateOrderRequest) Bind(r *http.Request) error {
	if err := req.Validate(r); err != nil {
		return err
	}
	return nil
}

func (req *CreateOrderRequest) Validate(r *http.Request) error {
	if err := validator.ValidateEmpty("product_id", req.ProductId); err != nil {
		return err
	}
	if err := validator.ValidateEmpty("qty", req.Qty); err != nil {
		return err
	}

	return nil
}

func (req *PaidOrderRequest) Bind(r *http.Request) error {
	if err := req.Validate(r); err != nil {
		return err
	}
	return nil
}

func (req *PaidOrderRequest) Validate(r *http.Request) error {
	if err := validator.ValidateEmpty("product_id", req.OrderId); err != nil {
		return err
	}

	return nil
}
