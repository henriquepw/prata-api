package balance

import (
	"net/http"

	"github.com/henriquepw/pobrin-api/internal/auth"
	"github.com/henriquepw/pobrin-api/pkg/httpx"
)

type balanceHandler struct {
	svc     BalanceService
	session auth.Session
}

func NewHandler(svc BalanceService, session auth.Session) *balanceHandler {
	return &balanceHandler{svc, session}
}

func (h *balanceHandler) PostUserBalance(w http.ResponseWriter, r *http.Request) {
	userID, err := h.session.GetUserID(r.Context())
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	body, err := httpx.GetBodyRequest[BalanceUpdate](r)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}
	body.UserID = userID

	item, err := h.svc.UpsertBalance(r.Context(), body)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	httpx.SuccessResponse(w, item)
}

func (h *balanceHandler) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	userID, err := h.session.GetUserID(r.Context())
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	item, err := h.svc.GetBalance(r.Context(), userID)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	httpx.SuccessResponse(w, item)
}
