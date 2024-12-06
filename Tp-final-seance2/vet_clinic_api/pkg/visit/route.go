package visit

import (

	"github.com/go-chi/chi/v5"
	"vet_clinic_api/config"
)

func Routes(cfg *config.Config) chi.Router {
	r := chi.NewRouter()

	r.Post("/", CreateVisit(cfg))          
	r.Get("/", GetAllVisits(cfg))         
	r.Get("/{id}", GetVisitByID(cfg))      
	r.Put("/{id}", UpdateVisit(cfg))       
	r.Delete("/{id}", DeleteVisit(cfg))    
	r.Get("/cat/{catID}", GetVisitsByCat(cfg)) 

	return r
}
