package transaction

import (
	"net/http"
	"time"

	"github.com/henriquepw/prata-api/internal/domains/auth"
	"github.com/henriquepw/prata-api/pkg/httpx"
	"github.com/henriquepw/prata-api/pkg/id"
)

type transactionHandler struct {
	svc TransactionService
}

func NewHandler(svc TransactionService) *transactionHandler {
	return &transactionHandler{svc}
}

func (h *transactionHandler) PostTransaction(w http.ResponseWriter, r *http.Request) {
	body, err := httpx.GetBodyRequest[[]TransactionCreate](r)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	userID := auth.GetUserID(r.Context())

	for i := range body {
		body[i].UserID = userID
	}

	trxs, err := h.svc.CreateTransactions(r.Context(), body)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	httpx.SuccessCreatedResponse(w, trxs)
}

func (h *transactionHandler) PatchTransactionByID(w http.ResponseWriter, r *http.Request) {
	id, err := id.Parse(r.PathValue("id"))
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	body, err := httpx.GetBodyRequest[TransactionUpdate](r)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	err = h.svc.UpdateTransaction(r.Context(), id, body)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	httpx.SuccessResponse(w)
}

func (h *transactionHandler) GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	id, err := id.Parse(r.PathValue("id"))
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	transaction, err := h.svc.GetTransaction(r.Context(), id)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	httpx.SuccessResponse(w, transaction)
}

func (h *transactionHandler) GetTransactionList(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserID(r.Context())

	q := r.URL.Query()
	query := TransactionQuery{
		UserID:        userID,
		Cursor:        q.Get("cursor"),
		Search:        q.Get("search"),
		Type:          TransactionType(q.Get("type")),
		ReceivedAtGte: httpx.GetQueryTime(q, "receivedAtGte"),
		ReceivedAtLte: httpx.GetQueryTime(q, "receivedAtLte"),
		Limit:         httpx.GetQueryInt(q, "limit", 10),
	}

	transaction := h.svc.ListTransaction(r.Context(), query)

	httpx.SuccessResponse(w, transaction)
}

func (h *transactionHandler) GetMonthlyTransactions(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserID(r.Context())

	now := time.Now()
	year, month, _ := now.Date()
	startOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)

	query := TransactionQuery{
		UserID:        userID,
		ReceivedAtGte: startOfMonth,
		ReceivedAtLte: endOfMonth,
		Limit:         0,
	}

	transaction := h.svc.ListTransaction(r.Context(), query)

	httpx.SuccessResponse(w, transaction.Items)
}

func (h *transactionHandler) DeleteTransactionByID(w http.ResponseWriter, r *http.Request) {
	id, err := id.Parse(r.PathValue("id"))
	if err != nil {
		return
	}

	err = h.svc.DeleteTransaction(r.Context(), id)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	httpx.SuccessResponse(w)
}
