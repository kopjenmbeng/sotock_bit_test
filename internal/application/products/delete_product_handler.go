package products

import (
	"net/http"

	"github.com/RoseRocket/xerrs"
	"github.com/go-chi/chi"
	"github.com/kopjenmbeng/sotock_bit_test/internal/middleware"
	"github.com/kopjenmbeng/sotock_bit_test/internal/utility/respond"
	"github.com/kopjenmbeng/sotock_bit_test/internal/utility/validator"
)

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	var (
		// err error
		rc  = r.Context()
		err error
	)
	ProductId := chi.URLParam(r, "product_id")
	if err := validator.ValidateEmpty("product_id", ProductId); err != nil {
		respond.Nay(w, r, http.StatusBadRequest, err)
		return
	}
	useCase := UseCaseFromContext(rc)
	code, err := useCase.Delete(rc, ProductId)
	if err != nil {
		middleware.GetLogEntry(r).Error(xerrs.Details(err, respond.ErrMaxStack))
		respond.Nay(w, r, code, err)
		return
	}
	respond.Yay(w, r, code, map[string]string{
		"message": "Data berhasil dihapus !",
	})
	return

}
