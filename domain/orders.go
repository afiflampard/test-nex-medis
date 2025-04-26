package domain

import (
	"time"

	"github.com/google/uuid"
)

const (
	OrderStatusPending   = "pending"
	OrderStatusPaid      = "paid"
	OrderStatusShipped   = "shipped"
	OrderStatusCompleted = "completed"
	OrderStatusCanceled  = "canceled"
)

type ResponseOfTopFive struct {
	UserID     uuid.UUID `gorm:"column:id" json:"id"`
	UserName   string    `gorm:"column:username" json:"username"`
	Email      string    `gorm:"column:email" json:"email"`
	TotalSpent float64   `gorm:"column:total_spent" json:"total_spent"`
}

type Order struct {
	ID          uuid.UUID  `gorm:"column:id" json:"id"`
	UserID      uuid.UUID  `gorm:"column:user_id" json:"user_id"`
	OrderDate   time.Time  `gorm:"column:order_date" json:"order_date"`
	CartID      uuid.UUID  `gorm:"column:cart_id" json:"cart_id"`
	Status      string     `gorm:"column:status" json:"status"`
	TotalAmount float64    `gorm:"column:total_amount" json:"total_amount"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`

	Cart       Cart        `gorm:"foreignKey:CartID" json:"cart,omitempty"`
	User       User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID;references:ID" json:"order_items"`
}

func (Order) TableName() string {
	return "orders"
}

func (o *Order) CreateNewOrder(cartID, userID uuid.UUID) {
	currentlyUpdated := time.Now()
	o.ID = uuid.New()
	o.UserID = userID
	o.OrderDate = time.Now()
	o.CartID = cartID
	o.Status = OrderStatusPending
	o.CreatedAt = time.Now()
	o.UpdatedAt = &currentlyUpdated
}

func (o *Order) UpdateStatusOrder(status string) {
	currentlyUpdated := time.Now()
	o.Status = status
	o.UpdatedAt = &currentlyUpdated
}
