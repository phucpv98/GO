package user

import "github.com/gin-gonic/gin"

type ProductRouter struct{}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	// Public Router
	productRouterPublic := Router.Group("/product")
	{
		productRouterPublic.GET("search")
		productRouterPublic.GET("/detail/:id")
	}

	// Private Router
}
