package auth

import (
	"context"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/henriquepw/prata-api/pkg/errorx"
	"github.com/henriquepw/prata-api/pkg/httpx"
	"github.com/henriquepw/prata-api/pkg/id"
	"github.com/henriquepw/prata-api/pkg/jwt"
)

type ctxKey string

const (
	ctxAuth = ctxKey("auth.clains")
)

func RequireAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		if len(bearer) < 7 {
			httpx.ErrorResponse(w, errorx.Unauthorized())
			return
		}

		token := bearer[7:]
		claims, err := jwt.Validade(token)
		if err != nil {
			log.Error(err.Error())
			httpx.ErrorResponse(w, errorx.Unauthorized())
			return
		}

		ctx := context.WithValue(r.Context(), ctxAuth, claims)
		request := r.WithContext(ctx)

		next.ServeHTTP(w, request)
	})
}

func GetUserID(ctx context.Context) id.ID {
	auth, ok := ctx.Value(ctxAuth).(*jwt.Claims)
	if !ok {
		return id.ID("")
	}

	return id.ID(auth.Subject)
}
