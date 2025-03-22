package recurrence

import (
	"net/http"

	"github.com/henriquepw/pobrin-api/pkg/httputil"
	"github.com/henriquepw/pobrin-api/pkg/id"
)

type recurrenceHandler struct {
	svc RecurrenceService
}

func NewHandler(svc RecurrenceService) *recurrenceHandler {
	return &recurrenceHandler{svc}
}

func (h *recurrenceHandler) PostRecurrence(w http.ResponseWriter, r *http.Request) {
	body, err := httputil.GetBodyRequest[RecurrenceCreate](r)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	recurrence, err := h.svc.CreateRecurrence(r.Context(), body)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	httputil.SuccessCreatedResponse(w, recurrence.ID.String())
}

func (h *recurrenceHandler) PatchRecurrenceByID(w http.ResponseWriter, r *http.Request) {
	id, err := id.Parse(r.PathValue("id"))
	if err != nil {
		return
	}

	body, err := httputil.GetBodyRequest[RecurrenceUpdate](r)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	err = h.svc.UpdateRecurrence(r.Context(), id, body)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	httputil.SuccessResponse(w)
}

func (h *recurrenceHandler) GetRecurrenceByID(w http.ResponseWriter, r *http.Request) {
	id, err := id.Parse(r.PathValue("id"))
	if err != nil {
		return
	}

	recurrence, err := h.svc.GetRecurrence(r.Context(), id)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	httputil.SuccessResponse(w, recurrence)
}

func (h *recurrenceHandler) GetRecurrenceList(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	query := RecurrenceQuery{
		Cursor: q.Get("cursor"),
		Limit:  httputil.GetQueryInt(q, "limit", 10),
	}

	recurrence := h.svc.ListRecurrence(r.Context(), query)

	httputil.SuccessResponse(w, recurrence)
}

func (h *recurrenceHandler) DeleteRecurrenceByID(w http.ResponseWriter, r *http.Request) {
	id, err := id.Parse(r.PathValue("id"))
	if err != nil {
		return
	}

	err = h.svc.DeleteRecurrence(r.Context(), id)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	httputil.SuccessResponse(w)
}
