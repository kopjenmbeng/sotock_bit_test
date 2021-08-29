package movies

import (
	"net/http"

	"github.com/RoseRocket/xerrs"
	"github.com/go-chi/chi"
	"github.com/kopjenmbeng/sotock_bit_test/internal/middleware"
	"github.com/kopjenmbeng/sotock_bit_test/internal/utility/respond"
	"github.com/kopjenmbeng/sotock_bit_test/internal/utility/validator"
)

func GetDetailHandler(w http.ResponseWriter, r *http.Request) {
	var (
		// err error
		rc = r.Context()
	)

	// app_code:=r.Header.Get("X-Client-id")
	Id := chi.URLParam(r, "id")
	if err := validator.ValidateEmpty("id", Id); err != nil {
		respond.Nay(w, r, http.StatusBadRequest, err)
		return
	}
	useCase := UseCaseFromContext(rc)
	data, code, err := useCase.GetDetail(rc, Id)
	if err != nil {
		middleware.GetLogEntry(r).Error(xerrs.Details(err, respond.ErrMaxStack))
		respond.Nay(w, r, code, err)
		return
	}
	respond.Yay(w, r, http.StatusOK, data)
	return
}
