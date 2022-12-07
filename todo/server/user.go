package server

import (
	"github.com/go-chi/chi/v5"
	"todo/handler"
)

func userRoutes(r chi.Router) {
	r.Group(func(user chi.Router) {
		user.Get("/info", handler.GetUserInfo)
		user.Delete("/logout", handler.Logout)
	})
}
