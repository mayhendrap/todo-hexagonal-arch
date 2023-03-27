package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-hexagonal-arch/internal/application/http/dto"
	"todo-hexagonal-arch/internal/core/domain"
	"todo-hexagonal-arch/internal/core/interfaces"
)

type authController struct {
	authService interfaces.AuthService
}

func NewAuthController(authService interfaces.AuthService) *authController {
	return &authController{
		authService: authService,
	}
}

func (ac *authController) Register(c *gin.Context) {
	var request dto.RegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	user := domain.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
	}

	token, err := ac.authService.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{"data": token})
}

func (ac *authController) Login(c *gin.Context) {
	var request dto.LoginRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	token, err := ac.authService.Login(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"data": token})
}
