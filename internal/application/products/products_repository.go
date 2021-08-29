package products

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/kopjenmbeng/sotock_bit_test/internal/dto"
)

type IProductsRepository interface {
	Create(ctx context.Context, prod dto.Products) (status int, err error)
	Update(ctx context.Context, prod dto.Products) (status int, err error)
	FindAll(ctx context.Context) (result []dto.Products, status int, err error)
	Delete(ctx context.Context, id string) (status int, err error)
}

type ProductsRepository struct {
	dbr sqlx.QueryerContext
	dbw *sqlx.DB
}

func NewProductsRepository(dbr sqlx.QueryerContext, dbw *sqlx.DB) IProductsRepository {
	return &ProductsRepository{dbr: dbr, dbw: dbw}
}

func (repo *ProductsRepository) Create(ctx context.Context, prod dto.Products) (status int, err error) {
	query := fmt.Sprintf(`
	INSERT INTO tbl_product (id,name,price,qty)
	VALUES(?,?,?,?);
	`)
	_, err = repo.dbw.ExecContext(ctx, query,
		&prod.Id,
		&prod.Name,
		&prod.Price,
		&prod.Qty,
	)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	status = http.StatusCreated
	return
}
func (repo *ProductsRepository) Update(ctx context.Context, prod dto.Products) (status int, err error) {
	query := fmt.Sprintf(`
	UPDATE tbl_product
	SET name = ?,
	price = ?,
	qty = ?
	WHERE id = ?
	`)
	_, err = repo.dbw.ExecContext(ctx, query,
		&prod.Name,
		&prod.Price,
		&prod.Qty,
		&prod.Id,
	)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	status = http.StatusOK
	return
}
func (repo *ProductsRepository) FindAll(ctx context.Context) (result []dto.Products, status int, err error) {
	query := fmt.Sprintf(`
		SELECT id,name,price,qty
		FROM tbl_product;
	`)
	rows, err := repo.dbr.QueryContext(ctx, query)
	defer rows.Close()
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	for rows.Next() {
		var prd dto.Products
		err = rows.Scan(&prd.Id, &prd.Name, &prd.Price, &prd.Qty)
		if err != nil {
			status = http.StatusInternalServerError
			return
		}
		result = append(result, prd)
	}
	status = http.StatusOK
	return
}
func (repo *ProductsRepository) Delete(ctx context.Context, id string) (status int, err error) {
	query := fmt.Sprintf(`
	DELETE FROM tbl_product
	where id=?
	`)
	_, err = repo.dbw.ExecContext(ctx, query, &id)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	status = http.StatusOK
	return
}
