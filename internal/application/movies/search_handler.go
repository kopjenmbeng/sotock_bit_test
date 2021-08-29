package movies

import (
	"net/http"
	"strconv"

	"github.com/RoseRocket/xerrs"
	"github.com/kopjenmbeng/sotock_bit_test/internal/middleware"
	"github.com/kopjenmbeng/sotock_bit_test/internal/utility/respond"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	var (
		// err error
		rc = r.Context()
	)

	// app_code:=r.Header.Get("X-Client-id")
	Search := r.URL.Query().Get("search")
	Page, err := strconv.ParseInt(r.URL.Query().Get("page"), 0, 64)
	useCase := UseCaseFromContext(rc)
	data, code, err := useCase.Search(rc, Search, int(Page))
	if err != nil {
		middleware.GetLogEntry(r).Error(xerrs.Details(err, respond.ErrMaxStack))
		respond.Nay(w, r, code, err)
		return
	}
	respond.Yay(w, r, http.StatusOK, data)
	return
}
