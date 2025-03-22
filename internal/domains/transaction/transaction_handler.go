package transaction

import (
	"net/http"

	"github.com/henriquepw/pobrin-api/pkg/httputil"
	"github.com/henriquepw/pobrin-api/pkg/id"
)

type transactionHandler struct {
	svc TransactionService
}

func NewHandler(svc TransactionService) *transactionHandler {
	return &transactionHandler{svc}
}

func (h *transactionHandler) PostTransaction(w http.ResponseWriter, r *http.Request) {
	body, err := httputil.GetBodyRequest[TransactionCreate](r)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	transaction, err := h.svc.CreateTransaction(r.Context(), body)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	httputil.SuccessCreatedResponse(w, transaction.ID.String())
}

func (h *transactionHandler) PatchTransactionByID(w http.ResponseWriter, r *http.Request) {
	id, err := id.Parse(r.PathValue("id"))
	if err != nil {
		return
	}

	body, err := httputil.GetBodyRequest[TransactionUpdate](r)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	err = h.svc.UpdateTransaction(r.Context(), id, body)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	httputil.SuccessResponse(w)
}

func (h *transactionHandler) GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	id, err := id.Parse(r.PathValue("id"))
	if err != nil {
		return
	}

	transaction, err := h.svc.GetTransaction(r.Context(), id)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	httputil.SuccessResponse(w, transaction)
}

func (h *transactionHandler) GetTransactionList(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	query := TransactionQuery{
		Cursor: q.Get("cursor"),
		Limit:  httputil.GetQueryInt(q, "limit", 10),
	}

	transaction := h.svc.ListTransaction(r.Context(), query)

	httputil.SuccessResponse(w, transaction)
}

func (h *transactionHandler) DeleteTransactionByID(w http.ResponseWriter, r *http.Request) {
	id, err := id.Parse(r.PathValue("id"))
	if err != nil {
		return
	}

	err = h.svc.DeleteTransaction(r.Context(), id)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	httputil.SuccessResponse(w)
}
