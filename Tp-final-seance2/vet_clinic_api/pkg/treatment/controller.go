package treatment

import (
	"net/http"
	"strconv"

	"vet_clinic_api/config"
	"vet_clinic_api/database/dbmodel"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func CreateTreatment(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var treatment dbmodel.Treatment
		if err := render.Bind(r, &treatment); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if err := cfg.TreatmentRepo.Create(&treatment); err != nil {
			http.Error(w, "Failed to create treatment", http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, treatment)
	}
}

func GetAllTreatments(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		treatments, err := cfg.TreatmentRepo.FindAll()
		if err != nil {
			http.Error(w, "Failed to fetch treatments", http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, treatments)
	}
}

func GetTreatmentByID(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		treatments, err := cfg.TreatmentRepo.FindAll() 
		if err != nil {
			http.Error(w, "Failed to fetch treatments", http.StatusInternalServerError)
			return
		}

		for _, treatment := range treatments {
			if treatment.ID == uint(id) {
				render.JSON(w, r, treatment)
				return
			}
		}

		http.Error(w, "Treatment not found", http.StatusNotFound)
	}
}

func UpdateTreatment(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var treatment dbmodel.Treatment
		if err := render.Bind(r, &treatment); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		treatment.ID = uint(id)

		if err := cfg.TreatmentRepo.Update(&treatment); err != nil {
			http.Error(w, "Failed to update treatment", http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, treatment)
	}
}

// DeleteTreatment supprime un traitement par son ID.
func DeleteTreatment(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		if err := cfg.TreatmentRepo.Delete(uint(id)); err != nil {
			http.Error(w, "Failed to delete treatment", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

// GetTreatmentsByVisit récupère les traitements associés à une visite.
func GetTreatmentsByVisit(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		visitID, err := strconv.Atoi(chi.URLParam(r, "visitID"))
		if err != nil {
			http.Error(w, "Invalid visit ID", http.StatusBadRequest)
			return
		}

		treatments, err := cfg.TreatmentRepo.FindByVisitID(uint(visitID))
		if err != nil {
			http.Error(w, "Failed to fetch treatments", http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, treatments)
	}
}
