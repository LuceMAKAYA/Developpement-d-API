package cat

import (
	"encoding/json"
	"net/http"
	"strconv"
	"vet_clinic_api/config"
	"vet_clinic_api/database/dbmodel"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func CreateCat(config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cat dbmodel.Cat
		if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := config.CatRepo.Create(&cat); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, cat)
	}
}

func GetAllCats(config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cats, err := config.CatRepo.FindAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, cats)
	}
}

func GetCatByID(config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		cat, err := config.CatRepo.FindByID(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, cat)
	}
}

func UpdateCat(config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		var cat dbmodel.Cat
		if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		cat.ID = uint(id)
		if err := config.CatRepo.Update(&cat); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, cat)
	}
}

func DeleteCat(config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		if err := config.CatRepo.Delete(uint(id)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
