package domain

import (
	"boilerplate/forms"
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderMutation interface {
	CreateCart(ctx context.Context, userID uuid.UUID) (*uuid.UUID, error)
	CreateCartItem(ctx context.Context, forms []forms.CartItemInput, userID uuid.UUID) (*uuid.UUID, error)
	FindCartByUserID(ctx context.Context, userID uuid.UUID) ([]Cart, error)
	Order(ctx context.Context, form forms.OrderItemInput, userID uuid.UUID) (*uuid.UUID, error)
	Checkout(ctx context.Context, form forms.CheckoutOrderInput) (*uuid.UUID, error)
	Shipping(ctx context.Context, idOrder uuid.UUID) (*uuid.UUID, error)
	Canceled(ctx context.Context, idOrder uuid.UUID) (*uuid.UUID, error)
	FindFiveTopClientAmount(ctx context.Context) ([]ResponseOfTopFive, error)

	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type gormMutationOrder struct {
	tx *gorm.DB
}

func NewGormMutationOrder(ctx context.Context, db *gorm.DB) OrderMutation {
	tx := db.WithContext(ctx).Begin()

	return &gormMutationOrder{
		tx: tx,
	}
}

func (gom *gormMutationOrder) CreateCart(ctx context.Context, userID uuid.UUID) (*uuid.UUID, error) {
	var (
		cart Cart
	)

	cart.CreateNewCart(uuid.UUID(userID))

	if err := gom.tx.Create(&cart).Error; err != nil {
		return nil, err
	}

	return &cart.ID, nil
}

func (gom *gormMutationOrder) CreateCartItem(ctx context.Context, forms []forms.CartItemInput, userID uuid.UUID) (*uuid.UUID, error) {
	var (
		cartItemList []CartItem
	)

	idCart, err := gom.CreateCart(ctx, userID)
	if err != nil {
		return nil, err
	}

	for _, cartItemInput := range forms {
		cartItem := CartItem{}
		cartItem.CreateNewCarItems(cartItemInput, *idCart)
		cartItemList = append(cartItemList, cartItem)
	}

	if err := gom.tx.Create(&cartItemList).Error; err != nil {
		return nil, err
	}

	return idCart, nil
}

func (gom *gormMutationOrder) FindCartByUserID(ctx context.Context, userID uuid.UUID) ([]Cart, error) {
	var (
		cartList []Cart
	)

	if err := gom.tx.Where("user_id = ?", userID).Preload("CartItems").Find(&cartList).Error; err != nil {
		return []Cart{}, err
	}

	return cartList, nil
}

func (gom *gormMutationOrder) Order(ctx context.Context, form forms.OrderItemInput, userID uuid.UUID) (*uuid.UUID, error) {
	var (
		order             Order
		cart              Cart
		cartItemList      []CartItem
		productIDListItem []uuid.UUID
		orderItemList     []OrderItem
		productMapInput   = make(map[uuid.UUID]int)
		productMap        = make(map[uuid.UUID]Products)
		productList       []Products
	)

	if err := gom.tx.First(&cart, form.CartID).Error; err != nil {
		return nil, err
	}

	for _, productOrder := range form.ProductorderList {
		productIDListItem = append(productIDListItem, productOrder.ProductID)
	}

	if err := gom.tx.Where("cart_id = ? AND product_id IN (?)", cart.ID, productIDListItem).Find(&cartItemList).Error; err != nil {
		return nil, err
	}
	if len(cartItemList) == 0 {
		return nil, fmt.Errorf("product not found")
	}

	if err := gom.tx.Where("id IN (?)", productIDListItem).Find(&productList).Error; err != nil {
		return nil, err
	}

	for _, product := range productList {
		productMap[product.ID] = product
	}

	order.CreateNewOrder(cart.ID, userID)

	for _, productOrder := range form.ProductorderList {
		productMapInput[productOrder.ProductID] += productOrder.Quantity

	}

	for key, value := range productMapInput {
		valueDB := productMap[key]
		if valueDB.Stock < value {
			return nil, fmt.Errorf("stock not enough")
		}
	}

	for productID, qty := range productMapInput {
		var orderItem OrderItem
		valueProduct := productMap[productID]
		price := valueProduct.Price * float64(qty)
		orderItem.CreateNewOrderItem(order.ID, productID, qty, price)
		orderItemList = append(orderItemList, orderItem)
	}

	if err := gom.tx.Create(&order).Error; err != nil {
		return nil, err
	}

	if err := gom.tx.Create(&orderItemList).Error; err != nil {
		return nil, err
	}

	return &order.ID, nil
}

func (gom *gormMutationOrder) Checkout(ctx context.Context, form forms.CheckoutOrderInput) (*uuid.UUID, error) {
	var (
		order              Order
		totalAmount        = 0.0
		productList        []Products
		productIDlist      []uuid.UUID
		productMapQuantity = make(map[uuid.UUID]int)
		productUpdate      []Products
	)

	if err := gom.tx.Preload("OrderItems").First(&order, form.OrderID).Error; err != nil {
		return nil, err
	}

	for _, value := range order.OrderItems {
		productIDlist = append(productIDlist, value.ProductID)
		productMapQuantity[value.ProductID] = value.Quantity
	}

	if err := gom.tx.Where("id IN (?)", productIDlist).Find(&productList).Error; err != nil {
		return nil, err
	}

	for _, product := range productList {
		valueProductCart := productMapQuantity[product.ID]
		product.Stock = product.Stock - valueProductCart
		productUpdate = append(productUpdate, product)
	}
	order.UpdateStatusOrder(OrderStatusPaid)

	for _, orderItems := range order.OrderItems {
		totalAmount += orderItems.Price
	}

	order.TotalAmount = totalAmount

	if err := gom.tx.Save(&order).Error; err != nil {
		return nil, err
	}

	if err := gom.tx.Save(&productUpdate).Error; err != nil {
		return nil, err
	}

	return &order.ID, nil
}

func (gom *gormMutationOrder) Shipping(ctx context.Context, idOrder uuid.UUID) (*uuid.UUID, error) {
	var (
		order Order
	)

	if err := gom.tx.First(&order, idOrder).Error; err != nil {
		return nil, err
	}
	order.UpdateStatusOrder(OrderStatusShipped)

	if err := gom.tx.Save(&order).Error; err != nil {
		return nil, err
	}

	return &order.ID, nil
}

func (gom *gormMutationOrder) Completed(ctx context.Context, idOrder uuid.UUID) (*uuid.UUID, error) {
	var (
		order Order
	)

	if err := gom.tx.First(&order, idOrder).Error; err != nil {
		return nil, err
	}
	order.UpdateStatusOrder(OrderStatusCompleted)

	if err := gom.tx.Save(&order).Error; err != nil {
		return nil, err
	}

	return &order.ID, nil
}

func (gom *gormMutationOrder) Canceled(ctx context.Context, idOrder uuid.UUID) (*uuid.UUID, error) {
	var (
		order                          Order
		productList, productUpdateList []Products
		productIDList                  []uuid.UUID
		mapOfProductItemsQty           = make(map[uuid.UUID]int)
	)

	if err := gom.tx.Preload("OrderItems").First(&order, idOrder).Error; err != nil {
		return nil, err
	}

	for _, productItem := range order.OrderItems {
		productIDList = append(productIDList, productItem.ProductID)
		mapOfProductItemsQty[productItem.ProductID] = productItem.Quantity
	}

	if err := gom.tx.Where("id IN (?)", productIDList).Find(&productList).Error; err != nil {
		return nil, err
	}

	if order.Status == OrderStatusCanceled {
		return nil, fmt.Errorf("status already canceled")
	}

	order.UpdateStatusOrder(OrderStatusCanceled)
	for _, product := range productList {
		value := mapOfProductItemsQty[product.ID]
		product.Stock = product.Stock + value
		productUpdateList = append(productUpdateList, product)
	}

	if err := gom.tx.Save(&order).Error; err != nil {
		return nil, err
	}

	if err := gom.tx.Save(&productUpdateList).Error; err != nil {
		return nil, err
	}

	return &order.ID, nil
}

func (g *gormMutationOrder) FindFiveTopClientAmount(ctx context.Context) ([]ResponseOfTopFive, error) {
	var (
		topFiveList []ResponseOfTopFive
	)

	rawDB := `WITH last_month_orders AS (
    SELECT
        user_id,
        SUM(total_amount) AS total_spent
    FROM
        orders
    WHERE
        order_date >= date_trunc('month', CURRENT_DATE) - INTERVAL '1 month'
        AND order_date < date_trunc('month', CURRENT_DATE)
        AND status IN ('paid', 'shipped', 'completed')
    GROUP BY
        user_id
	)
		SELECT
    	u.id,
    	u.username,
    	u.email,
    	lmo.total_spent
		FROM
    	last_month_orders lmo
	JOIN
    	users u ON u.id = lmo.user_id
	ORDER BY
    	lmo.total_spent DESC
	LIMIT 5;`

	if err := g.tx.Raw(rawDB).Find(&topFiveList).Error; err != nil {
		return nil, err
	}

	return topFiveList, nil
}

func (g *gormMutationOrder) Commit(ctx context.Context) error {
	return g.tx.Commit().Error
}

func (g *gormMutationOrder) Rollback(ctx context.Context) error {
	return g.tx.Rollback().Error
}
