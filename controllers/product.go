package controllers

import (
	"boilerplate/domain"
	"boilerplate/forms"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductServiceController struct {
	DB *gorm.DB
}

func NewProductServiceMutation(db *gorm.DB) *ProductServiceController {
	return &ProductServiceController{
		DB: db,
	}
}

// @Summary Create Product
// @Description Create a new product with the given details.
// @Tags Product
// @Accept  json
// @Produce  json
// @Param input body forms.ProductFormInput true "Product Input"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /product/create [post]
func (psc ProductServiceController) CreateProducts(c *gin.Context) {
	var (
		ctx                    = c.Request.Context()
		formInputCreateProduct forms.ProductFormInput
		userID                 = c.GetString("user_id")
	)

	if err := c.ShouldBindJSON(&formInputCreateProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err, "Data": nil})
		return
	}

	mutation := domain.NewGormMutationProduct(ctx, psc.DB)
	id, err := mutation.CreateProducts(ctx, formInputCreateProduct, uuid.MustParse(userID))
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

// @Summary Get Product by ID
// @Description Get a product by its ID.
// @Tags Product
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /product/{id} [get]
func (psc ProductServiceController) FindProductByID(c *gin.Context) {
	var (
		ctx       = c.Request.Context()
		idProduct = c.Param("id")
	)

	mutation := domain.NewGormMutationProduct(ctx, psc.DB)
	product, err := mutation.FindProductByID(ctx, uuid.MustParse(idProduct))
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
		"Data":    product,
	})
}

// @Summary Get List of Products
// @Description Get a list of products based on their status.
// @Tags Product
// @Accept  json
// @Produce  json
// @Param status body forms.ProductStatus true "Product Status"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /product [get]
func (psc ProductServiceController) FindProductList(c *gin.Context) {
	var (
		ctx        = c.Request.Context()
		statusList forms.ProductStatus
	)
	if err := c.ShouldBindJSON(&statusList); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err, "Data": nil})
		return
	}

	mutation := domain.NewGormMutationProduct(ctx, psc.DB)
	productList, err := mutation.FindProductList(ctx, statusList.Status)
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
		"Data":    productList,
	})
}
