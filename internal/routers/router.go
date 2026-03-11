package routers

// import (
// 	"go-ecommerce/internal/controller"
// 	"go-ecommerce/internal/middlewares"

// 	"github.com/gin-gonic/gin"
// )

// func NewRouter() *gin.Engine {
// 	r := gin.Default()
// 	r.Use((middlewares.AuthMiddleware()))

// 	v1 := r.Group("/v1")
// 	{
// 		v1.GET("/ping", controller.NewUserController().GetInfoUser)
// 		// v1.PUT("/ping", Pong)
// 		// v1.PATCH("/ping", Pong)
// 		// v1.DELETE("/ping", Pong)
// 		// v1.HEAD("/ping", Pong)
// 		// v1.OPTIONS("/ping", Pong)
// 	}
// 	return r
// }

// func Pong(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "pong",
// 	})
// }
