package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// Public Router
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}

	// Private Router
	adminRouterPrivate := Router.Group("/admin")
	{
		adminRouterPrivate.POST("/active_user")
	}
}
