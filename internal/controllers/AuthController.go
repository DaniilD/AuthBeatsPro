package controllers

import (
	"AuthBeatsPro/internal/operations"
	"encoding/json"
	"log"
	"net/http"
)

type AuthController struct {
	loginOperation *operations.LoginOperation
}

func NewAuthController(loginOperation *operations.LoginOperation) *AuthController {
	return &AuthController{
		loginOperation: loginOperation,
	}
}

func (controller *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response, err := controller.loginOperation.Handle(r)

	if err != nil {
		log.Println(err.Error())
	}

	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Println(err.Error())
	}
}
