package service

import (
	"context"
	"log"
	"net/http"
	"time"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func GracefullyShutDown(ctx context.Context) (err error) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
 	r.Use(middleware.RealIP)
 	r.Use(middleware.Logger)
 	r.Use(middleware.Recoverer)
 	r.Use(middleware.Timeout(60 * time.Second))


	r.Post("/api/v1/saveNote", SaveNote)
	r.Put("/api/v1/updateNote", UpdateNote)
	r.Delete("/api/v1/deleteNote", DeleteNote)
	r.Get("/api/v1/listAllSaved", ListAllSaved)
	r.Get("/api/v1/listAllArchived", ListAllArchived)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen:%s\n", err)
		}
	}()

	log.Printf("Server started on port 8080")

	<-ctx.Done()

	log.Printf("Server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("Server shutdown failed:%+s", err)
	}

	log.Printf("Server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}
