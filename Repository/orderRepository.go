package repository

import (
	"github.com/siddhant/shetkariBasketGo/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	GetByID(id int) (*models.Order, error)
	GetAll() ([]*models.Order, error)
	Insert(order *models.Order) error
	InsertOrderItem(OrderItems []*models.OrderItem) error
	Delete(id int) error
	DeleteOrderItem(id int) error
}

type GormOrderRepository struct {
	DB *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) *GormOrderRepository{
	return &GormOrderRepository{DB: db}
}

func (r *GormOrderRepository) GetByID(id int) (*models.Order, error) {
	var order models.Order
	if err := r.DB.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *GormOrderRepository) GetAll() ([]*models.Order, error) {
	var orders []*models.Order
	if err := r.DB.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *GormOrderRepository) InsertOrderItem(OrderItems []*models.OrderItem) error {
	return r.DB.Create(OrderItems).Error
}

func (r *GormOrderRepository) Insert(order *models.Order) error {
	return r.DB.Create(order).Error
}

func (r *GormOrderRepository) Delete(id int) error {
	return r.DB.Delete(&models.Order{},id).Error
}

func (r *GormOrderRepository) DeleteOrderItem(id int) error {
	return r.DB.Delete(&models.CartItem{},id).Error
}