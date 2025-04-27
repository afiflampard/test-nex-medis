package controllers

import (
	"boilerplate/domain"
	"boilerplate/forms"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderServiceController struct {
	DB *gorm.DB
}

func NewOrderServiceMutation(db *gorm.DB) *OrderServiceController {
	return &OrderServiceController{
		DB: db,
	}
}

// @Summary Create a new cart
// @Description Create a new cart for a user by their ID
// @Tags Order
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /order/create-cart [post]
func (osc OrderServiceController) CreateCart(c *gin.Context) {
	var (
		ctx    = c.Request.Context()
		userID = c.GetString("user_id")
	)
	mutation := domain.NewGormMutationOrder(ctx, osc.DB)

	id, err := mutation.CreateCart(ctx, uuid.MustParse(userID))
	if err != nil {
		mutation.Rollback(ctx)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}
	mutation.Commit(ctx)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully registered",
		"Data":    id,
	})
}

// @Summary Add items to cart
// @Description Add multiple items to a user's cart
// @Tags Order
// @Accept  json
// @Produce  json
// @Param body body []forms.CartItemInput true "Cart Item Input"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /order/create-cart-item [post]
func (osc OrderServiceController) CreateCartItem(c *gin.Context) {
	var (
		ctx           = c.Request.Context()
		cartItemInput []forms.CartItemInput
		userID        = c.GetString("user_id")
	)
	if err := c.ShouldBindJSON(&cartItemInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err, "Data": nil})
		return
	}

	mutation := domain.NewGormMutationOrder(ctx, osc.DB)
	id, err := mutation.CreateCartItem(ctx, cartItemInput, uuid.MustParse(userID))
	if err != nil {
		mutation.Rollback(ctx)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	mutation.Commit(ctx)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully registered",
		"Data":    id,
	})

}

// @Summary Find cart by user ID
// @Description Retrieve all carts for a specific user based on user ID
// @Tags Order
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /order/find-cart-by-user-id [get]
func (osc OrderServiceController) FindCartByUserID(c *gin.Context) {
	var (
		ctx    = c.Request.Context()
		userID = c.GetString("user_id")
	)

	mutation := domain.NewGormMutationOrder(ctx, osc.DB)
	cartList, err := mutation.FindCartByUserID(ctx, uuid.MustParse(userID))
	if err != nil {
		mutation.Rollback(ctx)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	mutation.Commit(ctx)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully registered",
		"Data":    cartList,
	})
}

// @Summary Place an order
// @Description Place an order for a user with the specified order details
// @Tags Order
// @Accept  json
// @Produce  json
// @Param body body forms.OrderItemInput true "Order Input"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /order/place [post]
func (osc OrderServiceController) Order(c *gin.Context) {
	var (
		ctx            = c.Request.Context()
		formInputOrder forms.OrderItemInput
		userID         = c.GetString("user_id")
	)

	if err := c.ShouldBindJSON(&formInputOrder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err, "Data": nil})
		return
	}

	mutation := domain.NewGormMutationOrder(ctx, osc.DB)
	id, err := mutation.Order(ctx, formInputOrder, uuid.MustParse(userID))
	if err != nil {
		mutation.Rollback(ctx)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	mutation.Commit(ctx)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully registered",
		"Data":    id,
	})
}

// @Summary Checkout order
// @Description Checkout and process the order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param body body forms.CheckoutOrderInput true "Checkout Input"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /order/checkout [post]
func (osc OrderServiceController) Checkout(c *gin.Context) {
	var (
		ctx               = c.Request.Context()
		formInputCheckout forms.CheckoutOrderInput
	)

	if err := c.ShouldBindJSON(&formInputCheckout); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err, "Data": nil})
		return
	}
	mutation := domain.NewGormMutationOrder(ctx, osc.DB)
	id, err := mutation.Checkout(ctx, formInputCheckout)
	if err != nil {
		mutation.Rollback(ctx)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}
	mutation.Commit(ctx)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully registered",
		"Data":    id,
	})
}

// @Summary Mark order as shipped
// @Description Update order status to "shipped"
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path string true "Order ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /order/shipping/{id} [post]
func (osc OrderServiceController) Shipping(c *gin.Context) {
	var (
		idOrder = c.Param("id")
		ctx     = c.Request.Context()
	)

	mutation := domain.NewGormMutationOrder(ctx, osc.DB)
	id, err := mutation.Shipping(ctx, uuid.MustParse(idOrder))
	if err != nil {
		mutation.Rollback(ctx)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}
	mutation.Commit(ctx)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully registered",
		"Data":    id,
	})
}

// @Summary Mark order as completed
// @Description Update order status to "completed"
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path string true "Order ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /order/completed/{id} [post]
func (osc OrderServiceController) Completed(c *gin.Context) {
	var (
		idOrder = c.Param("id")
		ctx     = c.Request.Context()
	)

	mutation := domain.NewGormMutationOrder(ctx, osc.DB)
	id, err := mutation.Completed(ctx, uuid.MustParse(idOrder))
	if err != nil {
		mutation.Rollback(ctx)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	mutation.Commit(ctx)
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully completed",
		"Data":    id,
	})
}

// @Summary Mark order as canceled
// @Description Update the status of an order to "canceled"
// @Tags Order
// @Accept  json
// @Produce  json
// @Param id path string true "Order ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /order/canceled/{id} [post]
func (osc OrderServiceController) Canceled(c *gin.Context) {
	var (
		idOrder = c.Param("id")
		ctx     = c.Request.Context()
	)

	mutation := domain.NewGormMutationOrder(ctx, osc.DB)
	id, err := mutation.Canceled(ctx, uuid.MustParse(idOrder))
	if err != nil {
		mutation.Rollback(ctx)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}
	mutation.Commit(ctx)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully registered",
		"Data":    id,
	})
}

// @Summary Retrieve the top 5 clients based on their total order amount
// @Description Get the top 5 clients who have spent the most on orders
// @Tags Order
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /order/top-clients [get]
func (osc OrderServiceController) FindFiveTopClientAmount(c *gin.Context) {
	var (
		ctx = c.Request.Context()
	)

	mutation := domain.NewGormMutationOrder(ctx, osc.DB)
	topFiveList, err := mutation.FindFiveTopClientAmount(ctx)
	if err != nil {
		mutation.Rollback(ctx)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}
	mutation.Commit(ctx)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully registered",
		"Data":    topFiveList,
	})
}
