package cat

import (
	"vet_clinic_api/config"

	"github.com/go-chi/chi/v5"
)

func Routes(config *config.Config) chi.Router {
	r := chi.NewRouter()
	r.Post("/", CreateCat(config))
	r.Get("/", GetAllCats(config))
	r.Get("/{id}", GetCatByID(config))
	r.Put("/{id}", UpdateCat(config))
	r.Delete("/{id}", DeleteCat(config))
	return r
}
