package repository

import (
	"github.com/siddhant/shetkariBasketGo/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByID(id int) (*models.User, error)
	GetWithCartDataById(id int) (*models.User, error)
	GetAll() ([]*models.User, error)
	Insert(user *models.User) error
	UpdateFieldsByID(id int, fields map[string]interface{}) error
	Delete(id int) error
}

type GormUserRepository struct {
	DB *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{DB: db}
}

func (r *GormUserRepository) GetByID(id int) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) GetWithCartDataById(id int) (*models.User, error) {
	var users models.User
	if err := r.DB.
		Where("id = ?", id).
		Preload("CartItems",func(db *gorm.DB) *gorm.DB {
			return db.Omit("Product") 
		}).
		Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *GormUserRepository) GetAll() ([] *models.User, error) {
	var users []*models.User
	if err := r.DB.First(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *GormUserRepository) Insert(user *models.User) error {
	return r.DB.Create(user).Error;
}

func (r *GormUserRepository) UpdateFieldsByID(id int, fields map[string]interface{}) error {
	return r.DB.Model(&models.User{}).Where("id = ?", id).Updates(fields).Error
}

func (r *GormUserRepository) Delete(id int) error {
	return r.DB.Delete(&models.User{},id).Error
}
