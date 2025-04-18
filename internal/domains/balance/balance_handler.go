package balance

import (
	"net/http"

	"github.com/henriquepw/pobrin-api/internal/auth"
	"github.com/henriquepw/pobrin-api/pkg/httpx"
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
		httpx.ErrorResponse(w, err)
		return
	}

	body, err := httpx.GetBodyRequest[BalanceCreate](r)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}
	body.UserID = session.Subject

	item, err := h.svc.CreateBalance(r.Context(), body)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	httpx.SuccessCreatedResponse(w, item)
}

func (h *balanceHandler) PutUserBalance(w http.ResponseWriter, r *http.Request) {
	session, err := auth.GetSession(r)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	body, err := httpx.GetBodyRequest[BalanceUpdate](r)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}
	body.UserID = session.Subject

	item, err := h.svc.UpdateBalance(r.Context(), body)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	httpx.SuccessResponse(w, item)
}

func (h *balanceHandler) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	session, err := auth.GetSession(r)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	item, err := h.svc.GetBalance(r.Context(), session.Subject)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	httpx.SuccessResponse(w, item)
}
