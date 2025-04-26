package domain

import (
	"time"

	"github.com/google/uuid"
)

type OrderItem struct {
	ID        uuid.UUID  `gorm:"column:id" json:"id"`
	OrderID   uuid.UUID  `gorm:"column:order_id" json:"order_id"`
	ProductID uuid.UUID  `gorm:"column:product_id" json:"product_id"`
	Quantity  int        `gorm:"column:quantity" json:"quantity"`
	Price     float64    `gorm:"column:price" json:"price"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`

	// Relasi opsional
	Order   Order    `gorm:"foreignKey:OrderID" json:"order,omitempty"`
	Product Products `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

// TableName sets the custom table name for GORM
func (OrderItem) TableName() string {
	return "order_items"
}

func (oi *OrderItem) CreateNewOrderItem(orderID, productID uuid.UUID, quantity int, price float64) {
	oi.ID = uuid.New()
	oi.OrderID = orderID
	oi.ProductID = productID
	oi.Quantity = quantity
	oi.Price = price
}
