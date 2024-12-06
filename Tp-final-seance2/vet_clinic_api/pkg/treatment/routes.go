package treatment

import (
	"vet_clinic_api/config"

	"github.com/go-chi/chi/v5"
)

func Routes(cfg *config.Config) chi.Router {
	r := chi.NewRouter()

	r.Post("/", CreateTreatment(cfg))                                         
	r.Get("/", GetAllTreatments(cfg))                                        
	r.Get("/{id}", GetTreatmentByID(cfg))                                     
	r.Put("/{id}", UpdateTreatment(cfg))                                      
	r.Delete("/{id}", DeleteTreatment(cfg))                                  

	return r
}
