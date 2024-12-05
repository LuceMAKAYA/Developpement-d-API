package database

import (
	"vet_clinic_api/database/dbmodel"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("vet_clinic.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = Migrate(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&dbmodel.Cat{}, &dbmodel.Visit{}, &dbmodel.Treatment{})
}
