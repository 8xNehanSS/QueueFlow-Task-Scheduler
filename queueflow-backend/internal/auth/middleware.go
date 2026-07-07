package auth

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type ContextKey string

const (
	UserID   ContextKey = "user_id"
	UserRole ContextKey = "role"
)

func AuthMiddleware(
	next http.Handler,
) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			cookie, err := r.Cookie("token")
			if err != nil {
				http.Error(
					w,
					"missing token",
					401,
				)
				return
			}

			if cookie.Value == "" {
				http.Error(
					w,
					"missing token",
					401,
				)
				return
			}

			token, err :=
				ValidateToken(
					cookie.Value,
				)

			if err != nil ||
				!token.Valid {

				http.Error(
					w,
					err.Error(),
					401,
				)

				return
			}

			claims :=
				token.Claims.(jwt.MapClaims)

			ctx :=
				context.WithValue(
					r.Context(),
					UserID,
					claims["user_id"],
				)

			ctx =
				context.WithValue(
					ctx,
					UserRole,
					claims["role"],
				)

			next.ServeHTTP(
				w,
				r.WithContext(ctx),
			)

		})
}

func Protected(
	handler http.Handler,
	permission string,
) http.Handler {

	return AuthMiddleware(
		RequirePermission(
			permission,
			handler,
		),
	)
}
