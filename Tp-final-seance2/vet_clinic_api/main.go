package main

import (
	"log"
	"net/http"
	"vet_clinic_api/config"
	"vet_clinic_api/pkg/cat"
	"vet_clinic_api/pkg/treatment"
	"vet_clinic_api/pkg/visit"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	
	r.Mount("/api/v1/cats", cat.Routes(cfg))
	r.Mount("/api/v1/visits", visit.Routes(cfg))
	r.Mount("/api/v1/treatments", treatment.Routes(cfg))

	log.Println("Serving on :8080")
	http.ListenAndServe(":8080", r)
}
