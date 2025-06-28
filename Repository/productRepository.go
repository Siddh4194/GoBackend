package repository

import (
	"github.com/siddhant/shetkariBasketGo/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetByID(id int) (*models.Product, error)
	GetAll() ([]*models.Product, error)
	Insert(product *models.Product) error
	UpdateFieldsByID(id int, fields map[string]interface{}) error
	Delete(id int) error
}

type GormProductRepository struct {
	DB *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) *GormProductRepository {
	return &GormProductRepository{DB: db}
}

func (r *GormProductRepository) GetByID(id int) (*models.Product, error) {
	var product models.Product
	if err := r.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *GormProductRepository) GetAll() ([] *models.Product, error) {
	var products []*models.Product
	if err := r.DB.First(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *GormProductRepository) Insert(product *models.Product) error {
	return r.DB.Create(product).Error;
}

func (r *GormProductRepository) UpdateFieldsByID(id int, fields map[string]interface{}) error {
	return r.DB.Model(&models.Product{}).Where("id = ?", id).Updates(fields).Error
}

func (r *GormProductRepository) Delete(id int) error {
	return r.DB.Delete(&models.Product{},id).Error
}
