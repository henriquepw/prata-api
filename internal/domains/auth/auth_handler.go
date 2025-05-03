package auth

import (
	"net/http"

	"github.com/henriquepw/prata-api/pkg/httpx"
)

type authHandler struct {
	svc AuthService
}

func NewHandler(svc AuthService) *authHandler {
	return &authHandler{svc}
}

func (h *authHandler) PostSignUp(w http.ResponseWriter, r *http.Request) {
	body, err := httpx.GetBodyRequest[SignUpData](r)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	data, err := h.svc.SignUp(r.Context(), body)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	httpx.SuccessResponse(w, data)
}

func (h *authHandler) PostSignIn(w http.ResponseWriter, r *http.Request) {
	body, err := httpx.GetBodyRequest[SignInData](r)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	data, err := h.svc.SignIn(r.Context(), body)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	httpx.SuccessResponse(w, data)
}

func (h *authHandler) PostRenew(w http.ResponseWriter, r *http.Request) {
	token := r.PathValue("token")

	data, err := h.svc.RefreshAccess(r.Context(), token)
	if err != nil {
		httpx.ErrorResponse(w, err)
		return
	}

	httpx.SuccessResponse(w, data)
}
