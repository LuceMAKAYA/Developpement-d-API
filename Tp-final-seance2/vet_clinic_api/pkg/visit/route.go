package visit

import (

	"github.com/go-chi/chi/v5"
	"vet_clinic_api/config"
)

func Routes(cfg *config.Config) chi.Router {
	r := chi.NewRouter()

	r.Post("/", CreateVisit(cfg))          // POST /visits
	r.Get("/", GetAllVisits(cfg))          // GET /visits
	r.Get("/{id}", GetVisitByID(cfg))      // GET /visits/{id}
	r.Put("/{id}", UpdateVisit(cfg))       // PUT /visits/{id}
	r.Delete("/{id}", DeleteVisit(cfg))    // DELETE /visits/{id}
	r.Get("/cat/{catID}", GetVisitsByCat(cfg)) // GET /visits/cat/{catID}

	return r
}
