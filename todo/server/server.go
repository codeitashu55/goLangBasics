package server

import (
	"context"
	"net/http"
	"time"
	"todo/handler"
	"todo/middlewares"
	"todo/models"
	"todo/utils"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	chi.Router
	server *http.Server
}

const (
	readTimeout       = 5 * time.Minute
	readHeaderTimeout = 30 * time.Second
	writeTimeout      = 5 * time.Minute
)

// SetupRoutes provides all the routes that can be used
func SetupRoutes() *Server {
	router := chi.NewRouter()
	router.Route("/v1", func(v1 chi.Router) {
		v1.Use(middlewares.CommonMiddlewares()...)
		v1.Post("/health", func(w http.ResponseWriter, r *http.Request) {
			utils.RespondJSON(w, http.StatusOK, struct {
				Status string `json:"status"`
			}{Status: "server is running"})
		})
		v1.Route("/", func(public chi.Router) {
			public.Post("/register", handler.RegisterUser)
			public.Post("/login", handler.LoginUser)
		})
		v1.Route("/user", func(user chi.Router) {
			user.Use(middlewares.AuthMiddleware)
			user.Group(userRoutes)
		})
		v1.Route("/task", func(task chi.Router) {
			task.Use(middlewares.AuthMiddleware)
		})
		v1.Route("/admin", func(admin chi.Router) {
			admin.Use(middlewares.AuthMiddleware)
			admin.Use(middlewares.ShouldHaveRole(models.RoleAdmin))
		})
	})
	return &Server{
		Router: router,
	}
}

func (svc *Server) Run(port string) error {
	svc.server = &http.Server{
		Addr:              port,
		Handler:           svc.Router,
		ReadTimeout:       readTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
	}
	return svc.server.ListenAndServe()
}

func (svc *Server) Shutdown(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return svc.server.Shutdown(ctx)
}
