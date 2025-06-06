package controllers

import (
	"boilerplate/domain"
	"boilerplate/forms"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserServiceController struct {
	DB *gorm.DB
}

var userForm = new(forms.UserForm)

func NewUserServiceMutation(db *gorm.DB) *UserServiceController {
	return &UserServiceController{
		DB: db,
	}
}

// Login godoc
// @Summary User login
// @Description Authenticate user and return JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body forms.LoginForm true "Login Form"
// @Success 200 {object} map[string]interface{} "Successfully logged in"
// @Failure 406 {object} map[string]interface{} "Invalid Login Details"
// @Router /auth/login [post]
func (ctrl UserServiceController) Login(c *gin.Context) {
	var (
		loginForm forms.LoginForm
		ctx       = c.Request.Context()
	)

	if validationErr := c.ShouldBindJSON(&loginForm); validationErr != nil {
		message := userForm.Login(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": message,
		})
		return
	}
	mutation := domain.NewGormMutationUser(ctx, ctrl.DB)
	user, token, err := mutation.Login(ctx, loginForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": "Invalid Login Details",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged in", "user": user, "token": token})
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags Auth
// @Accept json
// @Produce json
// @Param register body forms.RegisterForm true "Register Form"
// @Success 200 {object} map[string]interface{} "Successfully registered"
// @Failure 406 {object} map[string]interface{} "Validation Error"
// @Router /auth/register [post]
func (ctrl UserServiceController) Register(c *gin.Context) {
	var (
		registerForm forms.RegisterForm
		ctx          = c.Request.Context()
	)

	if err := c.ShouldBindJSON(&registerForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err,
		})
		return
	}
	mutation := domain.NewGormMutationUser(ctx, ctrl.DB)

	user, err := mutation.Register(ctx, registerForm)
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
		"Data":    user.ID,
	})
}

// @Summary Get User by ID
// @Description Get a user by their ID.
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /user/{id} [get]
func (ctrl UserServiceController) FindByID(c *gin.Context) {
	var (
		ctx    = c.Request.Context()
		userID = c.Param("id")
	)

	mutation := domain.NewGormMutationUser(ctx, ctrl.DB)
	user, err := mutation.FindByID(ctx, uuid.MustParse(userID))
	if err != nil {
		mutation.Rollback(ctx)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	mutation.Commit(ctx)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully Get User",
		"Data":    user,
	})

}

// @Summary Get User by Email
// @Description Get a user by their email.
// @Tags User
// @Accept  json
// @Produce  json
// @Param body body forms.FindByEmailForm true "User Email"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /user/find-by-email [post]
func (ctrl UserServiceController) FindByEmail(c *gin.Context) {
	var (
		ctx          = c.Request.Context()
		inputByEmail forms.FindByEmailForm
	)

	if err := c.ShouldBindJSON(&inputByEmail); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err,
		})
		return
	}
	mutation := domain.NewGormMutationUser(ctx, ctrl.DB)
	user, err := mutation.FindByEmail(ctx, inputByEmail.Email)
	if err != nil {
		mutation.Rollback(ctx)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	mutation.Commit(ctx)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully Get User",
		"Data":    user,
	})

}

// @Summary Get Users by Join Date
// @Description Get a list of users who joined on a specific date.
// @Tags User
// @Accept  json
// @Produce  json
// @Param body body forms.FindByJoinDateForm true "Join Date"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /user/find-by-join-date [post]
func (ctrl UserServiceController) FindByJoinDate(c *gin.Context) {
	var (
		ctx             = c.Request.Context()
		inputByJoinDate forms.FindByJoinDateForm
	)

	if err := c.ShouldBindJSON(&inputByJoinDate); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err,
		})
		return
	}
	mutation := domain.NewGormMutationUser(ctx, ctrl.DB)
	userList, err := mutation.FindByJoinDate(ctx, inputByJoinDate.JoinDate)
	if err != nil {
		mutation.Rollback(ctx)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	mutation.Commit(ctx)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully Get User",
		"Data":    userList,
	})

}
