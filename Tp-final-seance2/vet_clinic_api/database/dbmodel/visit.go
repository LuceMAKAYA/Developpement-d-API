package dbmodel

import (
	"errors"

	"gorm.io/gorm"
)

type Visit struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	CatID        uint   `json:"cat_id"`
	Date         string `json:"date"`
	Reason       string `json:"reason"`
	Veterinarian string `json:"veterinarian"`
}

type VisitRepository interface {
	Create(visit *Visit) error
	FindByID(id uint) (*Visit, error)
	FindByCatID(catID uint) ([]Visit, error)
	FindAll() ([]Visit, error)
	Update(visit *Visit) error
	Delete(id uint) error
}

type visitRepository struct {
	db *gorm.DB
}

func NewVisitRepository(db *gorm.DB) VisitRepository {
	return &visitRepository{db: db}
}

func (r *visitRepository) Create(visit *Visit) error {
	return r.db.Create(visit).Error
}

func (r *visitRepository) FindByID(id uint) (*Visit, error) {
	var visit Visit
	err := r.db.First(&visit, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &visit, err
}

func (r *visitRepository) FindByCatID(catID uint) ([]Visit, error) {
	var visits []Visit
	err := r.db.Where("cat_id = ?", catID).Find(&visits).Error
	return visits, err
}

func (r *visitRepository) FindAll() ([]Visit, error) {
	var visits []Visit
	err := r.db.Find(&visits).Error
	return visits, err
}

func (r *visitRepository) Update(visit *Visit) error {
	result := r.db.Model(&Visit{}).Where("id = ?", visit.ID).Updates(visit)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows updated, visit not found")
	}
	return nil
}

func (r *visitRepository) Delete(id uint) error {
	result := r.db.Delete(&Visit{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows deleted, visit not found")
	}
	return nil
}
