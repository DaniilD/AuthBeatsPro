package internal

import (
	"AuthBeatsPro/internal/controllers"
	"AuthBeatsPro/internal/operations"
	"github.com/gorilla/mux"
)

type RouterBuilder struct {
}

func NewRouterBuilder() *RouterBuilder {
	return &RouterBuilder{}
}

func (routerFactory *RouterBuilder) Build() *Router {
	return NewRouter(
		mux.NewRouter(),
		controllers.NewAuthController(operations.GetOperationLocator().GetLoginOperation()),
		controllers.NewHealthController(),
		controllers.NewUserController(),
	)
}
