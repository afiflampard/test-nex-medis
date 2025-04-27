package routes

import (
	"boilerplate/controllers"
	"boilerplate/db"
	"boilerplate/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizeRoleMiddleware(requiredRole []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		existingRole := false
		for _, roleRequired := range requiredRole {
			if role == roleRequired {
				existingRole = true
				break
			}
		}

		if !exists || !existingRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden: insufficient role"})
			return
		}
		c.Next()
	}
}

func Routes(router *gin.Engine) {

	v1 := router.Group("/v1")
	{
		user := controllers.NewUserServiceMutation(db.GetDB())

		v1.POST("/user/login", user.Login)
		v1.POST("/user/register", user.Register)
		v1.POST("/user/find-by-email", user.FindByEmail)
		v1.GET("/user/:id", middleware.AuthMiddleware(), user.FindByID)
		v1.POST("/user/find-by-join-date", middleware.AuthMiddleware(), AuthorizeRoleMiddleware([]string{"admin"}), user.FindByJoinDate)

		order := controllers.NewOrderServiceMutation(db.GetDB())
		orderRoute := v1.Group("/order")
		orderRoute.POST("/create", middleware.AuthMiddleware(), AuthorizeRoleMiddleware([]string{"client"}), order.CreateCart)
		orderRoute.POST("/create-cart-item", middleware.AuthMiddleware(), AuthorizeRoleMiddleware([]string{"client"}), order.CreateCartItem)
		orderRoute.GET("/cart/:id", middleware.AuthMiddleware(), AuthorizeRoleMiddleware([]string{"client"}), order.FindCartByUserID)
		orderRoute.POST("/order", middleware.AuthMiddleware(), AuthorizeRoleMiddleware([]string{"client"}), order.Order)
		orderRoute.POST("/checkout", middleware.AuthMiddleware(), AuthorizeRoleMiddleware([]string{"client"}), order.Checkout)
		orderRoute.POST("/shipping/:id", middleware.AuthMiddleware(), AuthorizeRoleMiddleware([]string{"seller"}), order.Shipping)
		orderRoute.POST("/cancelled/:id", middleware.AuthMiddleware(), AuthorizeRoleMiddleware([]string{"client"}), order.Canceled)
		orderRoute.POST("/completed/:id", middleware.AuthMiddleware(), AuthorizeRoleMiddleware([]string{"admin"}), order.Completed)
		orderRoute.POST("/get-five", middleware.AuthMiddleware(), AuthorizeRoleMiddleware([]string{"seller"}), order.FindFiveTopClientAmount)

		product := controllers.NewProductServiceMutation(db.GetDB())
		productRoute := v1.Group("/product")
		productRoute.POST("/create", middleware.AuthMiddleware(), AuthorizeRoleMiddleware([]string{"seller"}), product.CreateProducts)
		productRoute.GET("/:id", middleware.AuthMiddleware(), AuthorizeRoleMiddleware([]string{"client", "seller"}), product.FindProductByID)
		productRoute.GET("/", middleware.AuthMiddleware(), AuthorizeRoleMiddleware([]string{"client", "seller"}), product.FindProductList)
	}
}
