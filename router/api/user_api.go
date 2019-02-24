package api

import (
	"encoding/json"
	"net/http"

	"github.com/AMCCG/project-2-backend-golang/structure"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		getAllUsers(w, r)
	case http.MethodPost:
		addUsers(w, r)
	case http.MethodPut:
		updateUsers(w, r)
	case http.MethodDelete:
		deleteUsers(w, r)
	default:
		redirect(w, r)
	}
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	var response structure.ResponseApi
	response.Status = "success"
	response.Code = http.StatusOK
	json.NewEncoder(w).Encode(response)
}

func addUsers(w http.ResponseWriter, r *http.Request) {
	var response structure.ResponseApi
	response.Status = "create success"
	response.Code = http.StatusCreated
	json.NewEncoder(w).Encode(response)
}

func updateUsers(w http.ResponseWriter, r *http.Request) {
	var response structure.ResponseApi
	response.Status = "update success"
	response.Code = http.StatusOK
	json.NewEncoder(w).Encode(response)
}

func deleteUsers(w http.ResponseWriter, r *http.Request) {
	var response structure.ResponseApi
	response.Status = "delete success"
	response.Code = http.StatusOK
	json.NewEncoder(w).Encode(response)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	var response structure.ResponseApi
	response.Status = "error"
	response.Code = http.StatusNotFound
	json.NewEncoder(w).Encode(response)
}
