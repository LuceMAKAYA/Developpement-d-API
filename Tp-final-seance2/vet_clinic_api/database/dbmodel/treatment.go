package dbmodel

import (
	"errors"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

type Treatment struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	VisitID  uint   `json:"visit_id"`
	Medicine string `json:"medicine"`
	Dosage   string `json:"dosage"`
}

func (t *Treatment) Bind(r *http.Request) error {
	if strings.TrimSpace(t.Medicine) == "" {
		return errors.New("medicine is required")
	}
	if strings.TrimSpace(t.Dosage) == "" {
		return errors.New("dosage is required")
	}
	if t.VisitID == 0 {
		return errors.New("visit_id is required")
	}
	return nil
}

type TreatmentRepository interface {
	Create(treatment *Treatment) error
	FindAll() ([]Treatment, error)
	FindByVisitID(visitID uint) ([]Treatment, error)
	FindByID(id uint) (*Treatment, error)
	Update(treatment *Treatment) error
	Delete(id uint) error
}

type treatmentRepository struct {
	db *gorm.DB
}

func NewTreatmentRepository(db *gorm.DB) TreatmentRepository {
	return &treatmentRepository{db: db}
}

func (r *treatmentRepository) Create(treatment *Treatment) error {
	return r.db.Create(treatment).Error
}

func (r *treatmentRepository) FindAll() ([]Treatment, error) {
	var treatments []Treatment
	err := r.db.Find(&treatments).Error
	return treatments, err
}

func (r *treatmentRepository) FindByVisitID(visitID uint) ([]Treatment, error) {
	var treatments []Treatment
	err := r.db.Where("visit_id = ?", visitID).Find(&treatments).Error
	return treatments, err
}

func (r *treatmentRepository) FindByID(id uint) (*Treatment, error) {
	var treatment Treatment
	err := r.db.First(&treatment, id).Error
	return &treatment, err
}

func (r *treatmentRepository) Update(treatment *Treatment) error {
	return r.db.Save(treatment).Error
}

func (r *treatmentRepository) Delete(id uint) error {
	return r.db.Delete(&Treatment{}, id).Error
}
