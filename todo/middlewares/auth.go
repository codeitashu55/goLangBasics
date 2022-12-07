package middlewares

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"todo/database/dbHelper"
	"todo/models"
)

type ContextKeys string

const (
	userContext ContextKeys = "__userContext"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("x-api-key")
		user, err := dbHelper.GetUserBySession(apiKey)
		if err != nil || user == nil {
			logrus.WithError(err).Errorf("failed to get user with token: %s", apiKey)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), userContext, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UserContext(r *http.Request) *models.User {
	if user, ok := r.Context().Value(userContext).(*models.User); ok && user != nil {
		return user
	}
	return nil
}

func ShouldHaveRole(role models.Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := UserContext(r)
			if user == nil || len(user.Roles) == 0 {
				w.WriteHeader(http.StatusForbidden)
				return
			}
			for _, roleData := range user.Roles {
				if roleData.Role == role {
					next.ServeHTTP(w, r)
					return
				}
			}
			w.WriteHeader(http.StatusForbidden)
		})
	}
}
