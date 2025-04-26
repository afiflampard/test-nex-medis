package forms

import "github.com/google/uuid"

type OrderForm struct{}

type CartItemInput struct {
	ProductID uuid.UUID `form:"product_id" json:"product_id"`
	Quantity  int       `form:"quantity" json:"quantity"`
}

type OrderItemInput struct {
	CartID           uuid.UUID      `form:"cart_id" json:"cart_id"`
	ProductorderList []ProductOrder `form:"product_order_list" json:"product_order_list"`
}

type ProductOrder struct {
	ProductID uuid.UUID `form:"product_id" json:"product_id"`
	Quantity  int       `form:"quantity" json:"quantity"`
}

type CheckoutOrderInput struct {
	OrderID             uuid.UUID         `form:"order_id" json:"order_id"`
	ProductCheckoutList []ProductCheckout `form:"product_checkout_list" json:"product_checkout_list"`
}

type ProductCheckout struct {
	ProductID uuid.UUID `form:"product_id" json:"product_id"`
	Price     float64   `form:"price" json:"price"`
}
