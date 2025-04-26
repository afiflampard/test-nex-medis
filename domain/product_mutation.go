package domain

import (
	"boilerplate/forms"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductMutation interface {
	CreateProducts(ctx context.Context, form forms.ProductFormInput, userID uuid.UUID) (*uuid.UUID, error)
	FindProductByID(ctx context.Context, productID uuid.UUID) (*Products, error)
	FindProductList(ctx context.Context, status []string) (*[]Products, error)

	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type gormMutationProduct struct {
	tx *gorm.DB
}

func NewGormMutationProduct(ctx context.Context, db *gorm.DB) ProductMutation {
	tx := db.WithContext(ctx).Begin()

	return &gormMutationProduct{
		tx: tx,
	}
}

func (gp *gormMutationProduct) CreateProducts(ctx context.Context, form forms.ProductFormInput, userID uuid.UUID) (*uuid.UUID, error) {

	var (
		product Products
	)

	product.CreateNewProduct(form, uuid.UUID(userID))

	if err := gp.tx.Create(&product).Error; err != nil {
		return nil, err
	}

	return &product.ID, nil
}

func (gp *gormMutationProduct) FindProductByID(ctx context.Context, productID uuid.UUID) (*Products, error) {
	var product Products

	if err := gp.tx.Preload("User").First(&product, productID).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (gp *gormMutationProduct) FindProductList(ctx context.Context, status []string) (*[]Products, error) {
	var product []Products

	if err := gp.tx.Preload("User").Where("status IN (?)", status).Find(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (gp *gormMutationProduct) Commit(ctx context.Context) error {
	return gp.tx.Commit().Error
}

func (gp *gormMutationProduct) Rollback(ctx context.Context) error {
	return gp.tx.Rollback().Error
}
