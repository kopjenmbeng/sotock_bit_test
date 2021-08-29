package authentication

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/kopjenmbeng/sotock_bit_test/internal/dto"
)

type IAuthenticationRepository interface {
	GetUser(ctx context.Context, email string) (data *dto.User, status int, Err error)
	Register(ctx context.Context, customer dto.User) (status int, err error)
}

type AuthenticationRepository struct {
	dbr sqlx.QueryerContext
	dbw *sqlx.DB
}

func NewAuthenticationRepository(dbr sqlx.QueryerContext, dbw *sqlx.DB) IAuthenticationRepository {
	return &AuthenticationRepository{dbr: dbr, dbw: dbw}
}

func (repo *AuthenticationRepository) IsDuplicateEmail(ctx context.Context, email string) (bool, error) {
	query := fmt.Sprintf(`
	SELECT count(*) as total
	FROM tbl_user where email=? limit 1
	`)

	var total int = 0
	err := repo.dbw.QueryRowContext(ctx, query, &email).Scan(&total)
	if err != nil {
		return false, err
	}
	if total > 0 {
		return true, nil
	}

	return false, nil

}
func (repo *AuthenticationRepository) Register(ctx context.Context, customer dto.User) (status int, err error) {

	query := fmt.Sprintf(`
	INSERT INTO tbl_user(
		id, email, full_name, salt, password, iteration, security_length)
		VALUES (?, ?, ?, ?, ?, ?,?);
	`)

	isDuplicate, err := repo.IsDuplicateEmail(ctx, customer.Email)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if isDuplicate {
		return http.StatusBadRequest, errors.New("Email sudah terdaftar !.")
	}
	// usr:=dto.Customer{Id: uuid.New().String(),Email: email,FullName: full_name,Salt: salt,Password: password,Iteration: }
	_, err = repo.dbw.ExecContext(ctx, query,
		&customer.Id,
		&customer.Email,
		&customer.FullName,
		&customer.Salt,
		&customer.Password,
		&customer.Iteration,
		&customer.SecurityLength,
	)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}
	status = http.StatusCreated
	return
}
func (repo *AuthenticationRepository) GetUser(ctx context.Context, email string) (data *dto.User, status int, Err error) {
	query := fmt.Sprintf(`
	SELECT id, email, full_name, salt, password, iteration, security_length
	FROM tbl_user where email=? limit 1
	`)

	var cus dto.User
	err := repo.dbr.QueryRowxContext(ctx, query, &email).Scan(&cus.Id, &email, &cus.FullName, &cus.Salt, &cus.Password, &cus.Iteration, &cus.SecurityLength)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusUnauthorized, errors.New("Email belum terdaftar !")
		}
		return nil, http.StatusInternalServerError, err
	}
	return &cus, http.StatusOK, nil
}
