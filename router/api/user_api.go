package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/mongodb/mongo-go-driver/bson/primitive"

	"github.com/AMCCG/project-2-backend-golang/constant"
	"github.com/AMCCG/project-2-backend-golang/crud"
	"github.com/AMCCG/project-2-backend-golang/structure"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		get(w, r)
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

func get(w http.ResponseWriter, r *http.Request) {
	paramID := strings.TrimPrefix(r.URL.Path, constant.UsersURL)
	var users structure.User
	var response structure.ResponseApi
	if paramID != "" {
		id, err := primitive.ObjectIDFromHex(paramID)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Print(id)
		users.ID = id
		response = crud.FindByID(users)
	} else {
		response = crud.FindAll(users)
	}
	json.NewEncoder(w).Encode(response)
}

func add(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var users structure.User
	err := decoder.Decode(&users)
	if err != nil {
		panic(err)
	}
	var response structure.ResponseApi = crud.Insert(users)
	json.NewEncoder(w).Encode(response)
}

func update(w http.ResponseWriter, r *http.Request) {
	users := structure.User{FIRSTNAME: "Apisit", LASTNAME: "Amornchanchaigul"}
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
