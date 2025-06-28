package seeders

import (
	"time"

	"github.com/siddhant/shetkariBasketGo/models"
	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) error {
	if err := seedUsers(db); err != nil {
		return err
	}
	if err := seedCategories(db); err != nil {
		return err
	}
	if err := seedProducts(db); err != nil {
		return err
	}
	if err := seedAds(db); err != nil {
		return err
	}
	if err := seedCartItems(db); err != nil {
		return err
	}
	if err := seedOrders(db); err != nil {
		return err
	}
	if err := seedOrderItems(db); err != nil {
		return err
	}
	return nil
}

func seedUsers(db *gorm.DB) error {
	users := []models.User{
		{Name: "Siddhant Kadam", Phone: "9876543210", Email: "siddhant@example.com", Password: "hashed_pwd", Address: "Pune", CreatedAt: time.Now()},
		{Name: "Riya Sharma", Phone: "9123456789", Email: "riya@example.com", Password: "hashed_pwd", Address: "Mumbai", CreatedAt: time.Now()},
	}
	return db.Create(&users).Error
}


func seedCategories(db *gorm.DB) error {
	categories := []models.Category{
		{Name: "Vegetables", Description: "Fresh vegetables", CreatedAt: time.Now()},
		{Name: "Fruits", Description: "Organic fruits", CreatedAt: time.Now()},
	}
	return db.Create(&categories).Error
}

func seedProducts(db *gorm.DB) error {
	var categories []models.Category
	db.Find(&categories)

	products := []models.Product{
		{Name: "Tomato", Description: "Fresh red tomatoes", PricePerKG: 30.0, StockQuantityKG: 100, ImageURL: "https://via.placeholder.com/150", CategoryID: &categories[0].ID, CreatedAt: time.Now()},
		{Name: "Apple", Description: "Kashmiri apples", PricePerKG: 120.0, StockQuantityKG: 50, ImageURL: "https://via.placeholder.com/150", CategoryID: &categories[1].ID, CreatedAt: time.Now()},
	}
	return db.Create(&products).Error
}

func seedAds(db *gorm.DB) error {
	ads := []models.Ad{
		{Title: "Welcome Offer", ImageURL: "https://via.placeholder.com/300", Link: "/offers", Active: true, CreatedAt: time.Now()},
		{Title: "Summer Sale", ImageURL: "https://via.placeholder.com/300", Link: "/sale", Active: true, CreatedAt: time.Now()},
	}
	return db.Create(&ads).Error
}


func seedCartItems(db *gorm.DB) error {
	var users []models.User
	var products []models.Product
	db.Find(&users)
	db.Find(&products)

	items := []models.CartItem{
		{UserID: users[0].ID, ProductID: products[0].ID, QuantityKG: 2.5, AddedAt: time.Now()},
		{UserID: users[1].ID, ProductID: products[1].ID, QuantityKG: 1.0, AddedAt: time.Now()},
	}
	return db.Create(&items).Error
}

func seedOrders(db *gorm.DB) error {
	var users []models.User
	db.Find(&users)

	orders := []models.Order{
		{UserID: users[0].ID, TotalAmount: 75.0, Status: "PENDING", CreatedAt: time.Now()},
		{UserID: users[1].ID, TotalAmount: 120.0, Status: "COMPLETED", CreatedAt: time.Now()},
	}
	return db.Create(&orders).Error
}

func seedOrderItems(db *gorm.DB) error {
	var orders []models.Order
	var products []models.Product
	db.Find(&orders)
	db.Find(&products)

	orderItems := []models.OrderItem{
		{OrderID: orders[0].ID, ProductID: products[0].ID, QuantityKG: 2.5, PricePerKG: 30.0},
		{OrderID: orders[1].ID, ProductID: products[1].ID, QuantityKG: 1.0, PricePerKG: 120.0},
	}
	return db.Create(&orderItems).Error
}

