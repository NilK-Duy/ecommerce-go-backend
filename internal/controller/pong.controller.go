package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (p *PongController) Pong(c *gin.Context) {
	fmt.Println("--> My Handler")
	// Return JSON response
	name := c.DefaultQuery("name", "Kai")
	// c.ShouldBindJSON()
	uid := c.Query("uid") // http://localhost:8000/api/v1/ping\?uid\=1234
	c.JSON(http.StatusOK, gin.H{
		"message": "pong.hhh..ping " + name,
		"uid":     uid,
		"users":   []string{"Alice", "Bob", "Charlie"},
	})
}
