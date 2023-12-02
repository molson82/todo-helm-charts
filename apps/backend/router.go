package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(render.SetContentType(render.ContentTypeJSON),
		middleware.RedirectSlashes,
		middleware.RequestID,
		middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Mount("/heartbeat", heartbeatRoutes())
		r.Mount("/todo", todoRoutes())
	})

	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("X-Frame-Options", "DENY")
		rw.Header().Add("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		log.Println(rw.Write([]byte("online")))
	})

	return r
}
