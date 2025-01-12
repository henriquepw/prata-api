package income

import (
	"net/http"

	"github.com/henriquepw/pobrin-api/pkg/httputil"
)

type incomeHandler struct {
	svc IncomeService
}

func NewIncomeHandler(svc IncomeService) *incomeHandler {
	return &incomeHandler{svc}
}

func (h *incomeHandler) PostIncome(w http.ResponseWriter, r *http.Request) {
	body, err := httputil.GetBodyRequest[IncomeCreate](r)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	income, err := h.svc.CreateIncome(r.Context(), body)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	httputil.SuccessCreatedResponse(w, income.ID.String())
}

func (h *incomeHandler) PatchIncomeByID(w http.ResponseWriter, r *http.Request) {
	body, err := httputil.GetBodyRequest[IncomeUpdate](r)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	err = h.svc.UpdateIncome(r.Context(), r.PathValue("incomeId"), body)
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	httputil.SuccessResponse(w)
}

func (h *incomeHandler) GetIncomeByID(w http.ResponseWriter, r *http.Request) {
	income, err := h.svc.GetIncome(r.Context(), r.PathValue("incomeId"))
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	httputil.SuccessResponse(w, income)
}

func (h *incomeHandler) GetIncomeList(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	query := IncomeQuery{
		Cursor: q.Get("cursor"),
		Limit:  httputil.GetQueryInt(q, "limit", 10),
	}

	income := h.svc.ListIncome(r.Context(), query)

	httputil.SuccessResponse(w, income)
}

func (h *incomeHandler) DeleteIncomeByID(w http.ResponseWriter, r *http.Request) {
	err := h.svc.DeleteIncome(r.Context(), r.PathValue("incomeId"))
	if err != nil {
		httputil.ErrorResponse(w, err)
		return
	}

	httputil.SuccessResponse(w)
}
