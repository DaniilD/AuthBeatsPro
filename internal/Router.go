package internal

import (
	"AuthBeatsPro/internal/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	muxRouter        *mux.Router
	authController   *controllers.AuthController
	healthController *controllers.HealthController
	userController   *controllers.UserController
}

func NewRouter(router *mux.Router,
	authController *controllers.AuthController,
	healthController *controllers.HealthController,
	userController *controllers.UserController) *Router {

	return &Router{
		muxRouter:        router,
		authController:   authController,
		healthController: healthController,
		userController:   userController,
	}
}

func (router *Router) InitRouts() {
	router.muxRouter.HandleFunc("/user/login", router.authController.Login).Methods(http.MethodPost)

	router.muxRouter.HandleFunc("/get/user/{id}", router.userController.GetById).Methods(http.MethodGet)

	router.muxRouter.HandleFunc("/healthz/ping", router.healthController.Ping).Methods(http.MethodGet)
}
