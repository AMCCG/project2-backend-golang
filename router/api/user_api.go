package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/AMCCG/project-2-backend-golang/constant"
	"github.com/AMCCG/project-2-backend-golang/crud"
	"github.com/AMCCG/project-2-backend-golang/structure"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		id := strings.TrimPrefix(r.URL.Path, constant.UsersURL)
		if id != "" {
			get(w, r)
		} else {
			getAll(w, r)
		}
	case http.MethodPost:
		add(w, r)
	case http.MethodPut:
		update(w, r)
	case http.MethodDelete:
		delete(w, r)
	default:
		redirect(w, r)
	}
}

func getAll(w http.ResponseWriter, r *http.Request) {
	users := structure.User{FirstName: "Apisit", LastName: "Amornchanchaigul"}
	var response structure.ResponseApi = crud.FindAll(users)
	json.NewEncoder(w).Encode(response)
}

func get(w http.ResponseWriter, r *http.Request) {
	users := structure.User{FirstName: "Apisit", LastName: "AMCCG"}
	var response structure.ResponseApi
	response.Status = "success"
	response.Code = http.StatusOK
	response.Content = users
	json.NewEncoder(w).Encode(response)
}

func add(w http.ResponseWriter, r *http.Request) {
	users := structure.User{FirstName: "Apisit", LastName: "Amornchanchaigul"}
	var response structure.ResponseApi = crud.Insert(users)
	json.NewEncoder(w).Encode(response)
}

func update(w http.ResponseWriter, r *http.Request) {
	users := structure.User{FirstName: "Apisit", LastName: "Amornchanchaigul"}
	var response structure.ResponseApi = crud.Update(users)
	json.NewEncoder(w).Encode(response)
}

func delete(w http.ResponseWriter, r *http.Request) {
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
