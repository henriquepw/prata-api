package balance

import (
	"net/http"

	"github.com/henriquepw/pobrin-api/internal/auth"
	"github.com/henriquepw/pobrin-api/pkg/httputil"
)

type balanceHandler struct {
	svc BalanceService
}

func NewHandler(svc BalanceService) *balanceHandler {
	return &balanceHandler{svc}
}

func (h *balanceHandler) PostUserBalance(w http.ResponseWriter, r *http.Request) {
	session, err := auth.GetSession(r)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	body, err := httputil.GetBodyRequest[BalanceCreate](r)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}
	body.UserID = session.Subject

	item, err := h.svc.CreateBalance(r.Context(), body)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	httputil.SuccessCreatedResponse(w, item.ID.String())
}

func (h *balanceHandler) PutUserBalance(w http.ResponseWriter, r *http.Request) {
	session, err := auth.GetSession(r)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	body, err := httputil.GetBodyRequest[BalanceUpdate](r)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}
	body.UserID = session.Subject

	err = h.svc.UpdateBalance(r.Context(), body)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	httputil.SuccessResponse(w)
}

func (h *balanceHandler) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	session, err := auth.GetSession(r)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	item, err := h.svc.GetBalance(r.Context(), session.Subject)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	httputil.SuccessResponse(w, item)
}
