package products

import (
	"context"

	"github.com/google/uuid"
	"github.com/kopjenmbeng/sotock_bit_test/internal/dto"
)

type IProductsUsecase interface {
	Create(ctx context.Context, req AddProductRequest) (status int, err error)
	Update(ctx context.Context, req UpdateProductRequest) (status int, err error)
	GetAll(ctx context.Context) (result []dto.Products, status int, err error)
	Delete(ctx context.Context, id string) (status int, err error)
}

type ProductsUseCase struct {
	repository IProductsRepository
}

func NewProductsUseCase(repo IProductsRepository) IProductsUsecase {
	return &ProductsUseCase{repository: repo}
}

func (use_case *ProductsUseCase) Create(ctx context.Context, req AddProductRequest) (status int, err error) {
	prod := dto.Products{Id: uuid.New().String(), Name: req.Name, Price: req.Price, Qty: req.Qty}
	status, err = use_case.repository.Create(ctx, prod)
	return
}
func (use_case *ProductsUseCase) Update(ctx context.Context, req UpdateProductRequest) (status int, err error) {
	prod := dto.Products{Id: req.Id, Name: req.Name, Price: req.Price, Qty: req.Qty}
	status, err = use_case.repository.Update(ctx, prod)
	return
}
func (use_case *ProductsUseCase) GetAll(ctx context.Context) (result []dto.Products, status int, err error) {
	result, status, err = use_case.repository.FindAll(ctx)
	return
}
func (use_case *ProductsUseCase) Delete(ctx context.Context, id string) (status int, err error) {
	status, err = use_case.repository.Delete(ctx, id)
	return
}
