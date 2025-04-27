package domain

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        uuid.UUID  `gorm:"column:id" json:"id"`
	UserID    uuid.UUID  `gorm:"column:user_id" json:"user_id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`

	User     User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
	CartItem []CartItem `gorm:"foreignKey:CartID;references:ID" json:"cart_item,omitempty"`
}

func (c Cart) TableName() string {
	return "carts"
}

func (c *Cart) CreateNewCart(userID uuid.UUID) {
	currentlyUpdate := time.Now()
	c.ID = uuid.New()
	c.UserID = userID
	c.CreatedAt = time.Now()
	c.UpdatedAt = &currentlyUpdate
}
