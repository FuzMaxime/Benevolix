package authentification

import (
	"context"
	"net/http"
)

type contextKey string

const userIdContextKey contextKey = "userId"

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

			userId, err := ParseToken(secret, authHeader)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), userIdContextKey, userId)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserFromContext(ctx context.Context) uint {
	userId, ok := ctx.Value(userIdContextKey).(uint)
	if !ok {
		return 0
	}
	return userId
}
