package visit

import (
	"net/http"
	"strconv"

	"vet_clinic_api/config"
	"vet_clinic_api/database/dbmodel"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func CreateVisit(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var visit dbmodel.Visit
		if err := render.DecodeJSON(r.Body, &visit); err != nil { 
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if err := cfg.VisitRepo.Create(&visit); err != nil {
			http.Error(w, "Failed to create visit", http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, visit)
	}
}

func GetAllVisits(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		visits, err := cfg.VisitRepo.FindAll() 
		if err != nil {
			http.Error(w, "Failed to fetch visits", http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, visits)
	}
}

func GetVisitByID(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id")) 
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		visit, err := cfg.VisitRepo.FindByID(uint(id)) 
		if err != nil {
			http.Error(w, "Visit not found", http.StatusNotFound)
			return
		}
		render.JSON(w, r, visit)
	}
}

func UpdateVisit(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var visit dbmodel.Visit
		if err := render.DecodeJSON(r.Body, &visit); err != nil { 
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		visit.ID = uint(id)

		if err := cfg.VisitRepo.Update(&visit); err != nil {
			http.Error(w, "Failed to update visit", http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, visit)
	}
}

func DeleteVisit(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		if err := cfg.VisitRepo.Delete(uint(id)); err != nil { 
			http.Error(w, "Failed to delete visit", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func GetVisitsByCat(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		catID, err := strconv.Atoi(chi.URLParam(r, "catID"))
		if err != nil {
			http.Error(w, "Invalid cat ID", http.StatusBadRequest)
			return
		}

		visits, err := cfg.VisitRepo.FindByCatID(uint(catID))
		if err != nil {
			http.Error(w, "Failed to fetch visits", http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, visits)
	}
}
