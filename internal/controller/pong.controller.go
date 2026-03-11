package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (pc *PongController) Pong(c *gin.Context) {
	fmt.Println("pong")
	name := c.DefaultQuery("name", "world")
	c.JSON(200, gin.H{"message": "pong", "name": name})
}