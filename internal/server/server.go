package server

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/ory/graceful"
	"log"
	"net/http"
	"real-time-ranking/internal/app"
	"real-time-ranking/internal/config"
	"real-time-ranking/internal/db"
	"real-time-ranking/internal/handler"
	"real-time-ranking/internal/service"
	"time"
)

type Server struct {
	server *http.Server
}

func NewServer(cfg config.App) *Server {
	svc := service.NewService(db.DataSource())
	srv := handler.NewServer(svc)
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Route("/docs", func(r chi.Router) {
		fs := http.StripPrefix("/docs", http.FileServer(http.Dir("./docs")))
		r.Get("/*", fs.ServeHTTP)
	})

	handler.HandlerFromMux(srv, r)

	httpSrv := &http.Server{
		Handler:      r,
		Addr:         cfg.Port,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return &Server{server: httpSrv}
}

func (h *Server) Start() error {
	log.Printf("Server is listening on %s...\n", h.server.Addr)

	s := graceful.WithDefaults(h.server)

	shutdown := graceful.ShutdownFunc(func(ctx context.Context) error {
		log.Println("Shutting down gracefully...")
		app.Close()
		return s.Shutdown(context.Background())
	})

	return graceful.Graceful(s.ListenAndServe, shutdown)
}
