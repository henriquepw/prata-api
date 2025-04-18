package auth

import (
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/henriquepw/pobrin-api/pkg/errorx"
)

func GetSession(r *http.Request) (*clerk.SessionClaims, error) {
	claims, ok := clerk.SessionClaimsFromContext(r.Context())
	if !ok {
		return nil, errorx.Unauthorized()
	}

	return claims, nil
}
