package controller

import (
	"ecommerce-backend/internal/service"
	"ecommerce-backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// controller -> service -> repo -> models -> database
func (uc *UserController) GetUserByID(c *gin.Context) {
	// if err != nil {
	// 	response.ErrorResponse(c, 20003, "No need")
	// }
	response.SuccessResponse(c, 20001, []string{"User1", "User2", "User3"})
}
