package controllers

import "net/http"

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (controller *HealthController) Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}
