package authentification

import (
	"context"
	"net/http"
)

type contextKey string

const emailContextKey contextKey = "email"

type contextJWT struct {
	userId uint
}

func AuthMiddleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Missing token", http.StatusUnauthorized)
				return
			}

			email, err := ParseToken(secret, authHeader)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), emailContextKey, email)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserFromContext(ctx context.Context) string {
	email, ok := ctx.Value(emailContextKey).(string)
	if !ok {
		return ""
	}
	return email
}
