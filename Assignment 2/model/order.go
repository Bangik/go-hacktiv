package model

import "time"

type Order struct {
	OrderId      int       `gorm:"primaryKey" json:"orderId"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `gorm:"foreignKey:OrderId" json:"items"`
}
