package dbmodel

import (
	"gorm.io/gorm"
)

type Cat struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   int    `json:"age"`
}

type CatRepository interface {
	Create(cat *Cat) error
	FindAll() ([]Cat, error)
	FindByID(id uint) (*Cat, error)
	Update(cat *Cat) error
	Delete(id uint) error
}

type catRepository struct {
	db *gorm.DB
}

func NewCatRepository(db *gorm.DB) CatRepository {
	return &catRepository{db: db}
}

func (r *catRepository) Create(cat *Cat) error {
	return r.db.Create(cat).Error
}

func (r *catRepository) FindAll() ([]Cat, error) {
	var cats []Cat
	err := r.db.Find(&cats).Error
	return cats, err
}

func (r *catRepository) FindByID(id uint) (*Cat, error) {
	var cat Cat
	err := r.db.First(&cat, id).Error
	return &cat, err
}

func (r *catRepository) Update(cat *Cat) error {
	return r.db.Save(cat).Error
}

func (r *catRepository) Delete(id uint) error {
	return r.db.Delete(&Cat{}, id).Error
}
