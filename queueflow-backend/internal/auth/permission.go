package auth

import (
	"net/http"

	"queueflow/internal/permissions"
)

func RequirePermission(
	permission string,
	next http.Handler,
) http.Handler {

	return http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			roleValue := r.Context().Value(UserRole)

			if roleValue == nil {
				http.Error(
					w,
					"missing user role",
					http.StatusUnauthorized,
				)
				return
			}

			role, ok := roleValue.(string)

			if !ok {
				http.Error(
					w,
					"invalid user role",
					http.StatusUnauthorized,
				)
				return
			}

			if !permissions.HasPermission(
				role,
				permission,
			) {

				http.Error(
					w,
					"forbidden",
					403,
				)

				return
			}

			next.ServeHTTP(w, r)

		})
}
