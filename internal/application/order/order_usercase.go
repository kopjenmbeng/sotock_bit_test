package order

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/kopjenmbeng/sotock_bit_test/internal/dto"
	"github.com/kopjenmbeng/sotock_bit_test/internal/middleware/jwe_auth"
)

type IOrderUseCase interface {
	Create(ctx context.Context, req CreateOrderRequest) (status int, err error)
	UpdatePayment(ctx context.Context, req PaidOrderRequest) (status int, err error)
	GetMyOrder(ctx context.Context) (result []dto.MyOrder, status int, err error)
}

type OrderUseCase struct {
	repository IOrderRepository
	r          *http.Request
}

func NewOrderUserCase(repo IOrderRepository, r *http.Request) IOrderUseCase {
	return &OrderUseCase{repository: repo, r: r}
}
func (use_case *OrderUseCase) GetMyOrder(ctx context.Context) (result []dto.MyOrder, status int, err error) {
	claim := jwe_auth.GetClaims(use_case.r)
	result, status, err = use_case.repository.GetOrder(ctx, claim.Public.Subject)
	if err != nil {
		return
	}
	return
}
func (use_case *OrderUseCase) Create(ctx context.Context, req CreateOrderRequest) (status int, err error) {
	claim := jwe_auth.GetClaims(use_case.r)
	order := dto.Order{OrderId: uuid.New().String(), Status: "Pending", ProductId: req.ProductId, UserId: claim.Public.Subject, Qty: req.Qty}
	status, err = use_case.repository.Create(ctx, order)
	if err != nil {
		return
	}
	return

}

func (use_case *OrderUseCase) UpdatePayment(ctx context.Context, req PaidOrderRequest) (status int, err error) {
	claim := jwe_auth.GetClaims(use_case.r)
	order := dto.Order{OrderId: req.OrderId, Status: "Paid", UserId: claim.Public.Subject}
	status, err = use_case.repository.UpdatePayment(ctx, order)
	if err != nil {
		return
	}
	return
}
