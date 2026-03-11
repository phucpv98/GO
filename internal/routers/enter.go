package routers

import (
	"go-ecommerce/internal/routers/manage"
	"go-ecommerce/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.UserRouterGroup
}

var RouterGroupApp = new(RouterGroup)
