package treatment

import (
	"vet_clinic_api/config"

	"github.com/go-chi/chi/v5"
)

func Routes(cfg *config.Config) chi.Router {
	r := chi.NewRouter()

	r.Post("/", CreateTreatment(cfg))                                         // POST /treatments
	r.Get("/", GetAllTreatments(cfg))                                         // GET /treatments
	r.Get("/{id}", GetTreatmentByID(cfg))                                     // GET /treatments/{id}
	r.Put("/{id}", UpdateTreatment(cfg))                                      // PUT /treatments/{id}
	r.Delete("/{id}", DeleteTreatment(cfg))                                   // DELETE /treatments/{id}

	return r
}
