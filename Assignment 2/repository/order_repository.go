package repository

import (
	"hacktiv-assignment-2/model"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order model.Order) error
	FindAll() ([]model.Order, error)
	FindByID(orderID int) (model.Order, error)
	Update(order model.Order) error
	Delete(orderID int) error
}

type orderRepository struct {
	db *gorm.DB
}

func (o *orderRepository) Create(order model.Order) error {
	tx := o.db.Begin()
	err := tx.Create(&order).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (o *orderRepository) FindAll() ([]model.Order, error) {
	var orders []model.Order
	err := o.db.Preload("Items").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *orderRepository) FindByID(orderID int) (model.Order, error) {
	var order model.Order
	err := o.db.Preload("Items").First(&order, orderID).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (o *orderRepository) Update(order model.Order) error {
	tx := o.db.Begin()
	err := tx.Save(&order).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (o *orderRepository) Delete(orderID int) error {
	err := o.db.Delete(&model.Order{}, orderID).Error
	if err != nil {
		return err
	}

	return nil
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}
