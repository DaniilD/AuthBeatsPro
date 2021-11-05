package controllers

import "net/http"

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (controller *UserController) GetById(w http.ResponseWriter, r *http.Request) {

}
