package order

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/kopjenmbeng/sotock_bit_test/internal/dto"
)

type IOrderRepository interface {
	Create(ctx context.Context, odr dto.Order) (status int, err error)
	UpdatePayment(ctx context.Context, order dto.Order) (status int, err error)
	GetOrder(ctx context.Context, user_id string) (result []dto.MyOrder, status int, err error)
}

type OrderRepository struct {
	dbr sqlx.QueryerContext
	dbw *sqlx.DB
}

func NewOrderRepository(dbr sqlx.QueryerContext, dbw *sqlx.DB) IOrderRepository {
	return &OrderRepository{dbr: dbr, dbw: dbw}
}
func (repo *OrderRepository) GetOrder(ctx context.Context, user_id string) (result []dto.MyOrder, status int, err error) {
	query := fmt.Sprintf(`
	SELECT order_id,product_id,b.name as product,a.qty,amount,status
	FROM tbl_order a inner join
	tbl_product b on b.id=a.product_id
	where a.user_id=?
	`)

	rows, err := repo.dbr.QueryContext(ctx, query, &user_id)
	defer rows.Close()
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	for rows.Next() {
		var order dto.MyOrder
		err = rows.Scan(&order.OrderId, &order.ProductId, &order.ProductName, &order.Qty, &order.Amount, &order.Status)
		if err != nil {
			status = http.StatusInternalServerError
			return
		}
		result = append(result, order)
	}
	status = http.StatusOK
	return
}

func (repo *OrderRepository) UpdatePayment(ctx context.Context, order dto.Order) (status int, err error) {
	query := fmt.Sprintf(`
	UPDATE tbl_order
	set status=?
	where order_id=? and user_id=?
	`)
	save, err := repo.dbw.ExecContext(ctx, query, &order.Status, &order.OrderId, order.UserId)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	affect, err := save.RowsAffected()
	if affect == 0 {
		return http.StatusBadRequest, errors.New("tidak ada pembayaran yang di update !")
	}
	status = http.StatusOK
	return

}
func (repo *OrderRepository) Create(ctx context.Context, odr dto.Order) (status int, Err error) {

	query := fmt.Sprintf(`
	INSERT INTO tbl_order(
		order_id,user_id,product_id,qty,amount, status)
		VALUES (?, ?, ?, ?,?,?);
	`)
	tx, err := repo.dbw.BeginTx(ctx, nil)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// check stock
	product, err := repo.GetProductById(ctx, odr.ProductId, tx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	if product.Qty < odr.Qty {
		return http.StatusBadRequest, errors.New("Stock tidak cukup !")
	}

	// create order
	odr.Amount = product.Price * float64(odr.Qty)
	_, err = tx.ExecContext(ctx, query,
		&odr.OrderId,
		&odr.UserId,
		&odr.ProductId,
		&odr.Qty,
		&odr.Amount,
		&odr.Status,
	)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// update stock
	status, err = repo.UpdateStock(ctx, odr.ProductId, product.Qty-odr.Qty, tx)
	if err != nil {
		tx.Rollback()
		return status, err
	}

	// commit all transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, err
	}
	return
}

func (repo *OrderRepository) GetProductById(ctx context.Context, product_id string, tx *sql.Tx) (*dto.Products, error) {
	query := fmt.Sprintf(`
	select id,name,price,qty
	from tbl_product  where id=?
	`)
	var product dto.Products

	err := tx.QueryRowContext(ctx, query, &product_id).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Qty,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product tidak ditemukan")
		}
		return nil, err
	}

	return &product, nil
}
func (repo *OrderRepository) UpdateStock(ctx context.Context, product_id string, remain int, tx *sql.Tx) (status int, err error) {
	query := fmt.Sprintf(`
	UPDATE tbl_product 
	SET qty=?
	WHERE id=?
	`)
	_, err = tx.ExecContext(ctx, query,
		&remain,
		&product_id,
	)
	if err != nil {
		return http.StatusInternalServerError, errors.New(fmt.Sprintf("error at step update stock %s", err.Error()))
	}
	return http.StatusCreated, nil
}
