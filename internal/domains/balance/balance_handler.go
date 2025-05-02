package balance

import (
	"net/http"

	"github.com/henriquepw/prata-api/internal/domains/auth"
	"github.com/henriquepw/prata-api/pkg/httpx"
)

type balanceHandler struct {
	svc BalanceService
}

func NewHandler(svc BalanceService) *balanceHandler {
	return &balanceHandler{svc}
}

func (h *balanceHandler) PostUserBalance(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.GetUserID(r)
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
	userID, err := auth.GetUserID(r)
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
