package domain

import (
	"boilerplate/forms"
	"time"

	"github.com/google/uuid"
)

type CartItem struct {
	ID        uuid.UUID  `gorm:"column:id" json:"id"`
	CartID    uuid.UUID  `gorm:"column:cart_id" json:"cart_id"`
	ProductID uuid.UUID  `gorm:"column:product_id" json:"product_id"`
	Quantity  int        `gorm:"column:quantity" json:"quantity"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`

	Cart    Cart     `gorm:"foreignKey:CartID" json:"cart,omitempty"`
	Product Products `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

func (ci CartItem) TableName() string {
	return "cart_items"
}

func (ci *CartItem) CreateNewCarItems(form forms.CartItemInput, cartID uuid.UUID) {
	currentlyUpdate := time.Now()

	ci.ID = uuid.New()
	ci.CartID = cartID
	ci.ProductID = form.ProductID
	ci.Quantity = form.Quantity
	ci.CreatedAt = time.Now()
	ci.UpdatedAt = &currentlyUpdate
}
