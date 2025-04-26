package domain

import (
	"boilerplate/forms"
	"time"

	"github.com/google/uuid"
)

type Products struct {
	ID          uuid.UUID  `gorm:"column:id" json:"id"`
	Name        string     `gorm:"column:name" json:"name"`
	Description string     `gorm:"column:description" json:"description"`
	Price       float64    `gorm:"column:price" json:"price"`
	Stock       int        `gorm:"column:stock" json:"stock"`
	Status      string     `gorm:"column:status" json:"status"`
	UserID      uuid.UUID  `gorm:"column:user_id" json:"user_id"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`

	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (p Products) TableName() string {
	return "products"
}

func (p *Products) CreateNewProduct(form forms.ProductFormInput, userID uuid.UUID) {
	newCurrently := time.Now()
	p.ID = uuid.New()
	p.Name = form.Name
	p.Description = form.Description
	p.Price = form.Price
	p.Stock = form.Stock
	p.Status = form.Status
	p.UserID = userID
	p.CreatedAt = time.Now()
	p.UpdatedAt = &newCurrently
}

func (p *Products) UpdateProductStatus(status string) {
	newCurrently := time.Now()
	p.Status = status
	p.UpdatedAt = &newCurrently
}
