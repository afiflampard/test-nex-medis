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
