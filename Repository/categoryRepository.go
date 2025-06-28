package repository

import (
	"github.com/siddhant/shetkariBasketGo/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetByID(id int) (*models.Category, error)
	GetAll() ([]*models.Category, error)
	GetAllWithProduct() ([]*models.Category, error)
	Insert(category *models.Category) error
	UpdateFieldsByID(id int, fields map[string]interface{}) error
	Delete(id int) error
}

type GormCategoryRepository struct {
	DB *gorm.DB
}

func NewGormCategoryRepository(db *gorm.DB) *GormCategoryRepository {
	return &GormCategoryRepository{DB: db}
}


func (r *GormCategoryRepository) GetByID(id int) (*models.Category, error) {
	var category models.Category
	if err := r.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *GormCategoryRepository) GetAll() ([] *models.Category, error) {
	var categories []*models.Category
	if err := r.DB.First(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}


func (r *GormCategoryRepository) GetAllWithProduct() ([] *models.Category, error) {
	var categories []*models.Category
	if err := r.DB.Preload("Product").First(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *GormCategoryRepository) Insert(category *models.Category) error {
	return r.DB.Create(category).Error;
}

func (r *GormCategoryRepository) UpdateFieldsByID(id int, fields map[string]interface{}) error {
	return r.DB.Model(&models.Category{}).Where("id = ?", id).Updates(fields).Error
}

func (r *GormCategoryRepository) Delete(id int) error {
	return r.DB.Delete(&models.Category{},id).Error
}
