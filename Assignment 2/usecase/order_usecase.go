package usecase

import (
	"hacktiv-assignment-2/model"
	"hacktiv-assignment-2/repository"
)

type OrderUsecase interface {
	Create(order model.Order) error
	FindAll() ([]model.Order, error)
	FindByID(orderID int) (model.Order, error)
	Update(order model.Order) error
	Delete(orderID int) error
}

type orderUsecase struct {
	repository repository.OrderRepository
}

func (o *orderUsecase) Create(order model.Order) error {
	return o.repository.Create(order)
}

func (o *orderUsecase) FindAll() ([]model.Order, error) {
	return o.repository.FindAll()
}

func (o *orderUsecase) FindByID(orderID int) (model.Order, error) {
	return o.repository.FindByID(orderID)
}

func (o *orderUsecase) Update(order model.Order) error {
	return o.repository.Update(order)
}

func (o *orderUsecase) Delete(orderID int) error {
	return o.repository.Delete(orderID)
}

func NewOrderUsecase(repository repository.OrderRepository) OrderUsecase {
	return &orderUsecase{repository}
}
