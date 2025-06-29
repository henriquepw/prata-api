package recurrence

import (
	"net/http"

	"github.com/henriquepw/prata-api/internal/domains/auth"
	"github.com/henriquepw/prata-api/internal/domains/transaction"
	"github.com/henriquepw/prata-api/pkg/httpx"
	"github.com/henriquepw/prata-api/pkg/id"
)

type recurrenceHandler struct {
	svc RecurrenceService
}

func NewHandler(svc RecurrenceService) *recurrenceHandler {
	return &recurrenceHandler{svc}
}

func (h *recurrenceHandler) PostRecurrence(w http.ResponseWriter, r *http.Request) {
	body, err := httpx.GetBodyRequest[RecurrenceCreate](r)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	body.UserID = auth.GetUserID(r.Context())

	recurrence, err := h.svc.CreateRecurrence(r.Context(), body)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	httpx.SuccessCreatedResponse(w, recurrence.ID.String())
}

func (h *recurrenceHandler) PatchRecurrenceByID(w http.ResponseWriter, r *http.Request) {
	id, err := id.Parse(r.PathValue("id"))
	if err != nil {
		return
	}

	body, err := httpx.GetBodyRequest[RecurrenceUpdate](r)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	err = h.svc.UpdateRecurrence(r.Context(), id, body)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	httpx.SuccessResponse(w)
}

func (h *recurrenceHandler) GetRecurrenceByID(w http.ResponseWriter, r *http.Request) {
	id, err := id.Parse(r.PathValue("id"))
	if err != nil {
		return
	}

	recurrence, err := h.svc.GetRecurrence(r.Context(), id)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	httpx.SuccessResponse(w, recurrence)
}

func (h *recurrenceHandler) GetRecurrenceList(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserID(r.Context())

	q := r.URL.Query()
	query := RecurrenceQuery{
		UserID:     userID,
		Cursor:     q.Get("cursor"),
		Search:     q.Get("search"),
		Frequence:  Frequence(q.Get("frequence")),
		Type:       transaction.TransactionType(q.Get("type")),
		StartAtGte: httpx.GetQueryTime(q, "startAtGte"),
		StartAtLte: httpx.GetQueryTime(q, "startAtLte"),
		EndAtGte:   httpx.GetQueryTime(q, "endAtGte"),
		EndAtLte:   httpx.GetQueryTime(q, "endAtLte"),
		Limit:      httpx.GetQueryInt(q, "limit", 10),
	}

	recurrence := h.svc.ListRecurrence(r.Context(), query)

	httpx.SuccessResponse(w, recurrence)
}

func (h *recurrenceHandler) DeleteRecurrenceByID(w http.ResponseWriter, r *http.Request) {
	id, err := id.Parse(r.PathValue("id"))
	if err != nil {
		return
	}

	err = h.svc.DeleteRecurrence(r.Context(), id)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	httpx.SuccessResponse(w)
}
