package config

import (
	"vet_clinic_api/database"
	"vet_clinic_api/database/dbmodel"

	"gorm.io/gorm"
)

type Config struct {
	DB            *gorm.DB
	CatRepo       dbmodel.CatRepository
	VisitRepo     dbmodel.VisitRepository
	TreatmentRepo dbmodel.TreatmentRepository

}

func New() (*Config, error) {
	db, err := database.InitDB()
	if err != nil {
		return nil, err
	}

	config := &Config{
		DB:            db,
		CatRepo:       dbmodel.NewCatRepository(db),
		VisitRepo:     dbmodel.NewVisitRepository(db),
		TreatmentRepo: dbmodel.NewTreatmentRepository(db),

	}

	return config, nil
}
