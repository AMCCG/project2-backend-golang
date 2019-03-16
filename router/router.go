package router

import (
	"encoding/json"
	"net/http"

	"project-2-backend-golang/constant"
	"project-2-backend-golang/router/api"
	"project-2-backend-golang/structure"
)

func Initial() {
	http.HandleFunc(constant.UsersURL, api.UserHandler)
	http.HandleFunc("/api/v1/login", login)
	http.ListenAndServe(":"+constant.PORT, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	var response structure.ResponseApi
	response.Status = "Login success"
	response.Code = http.StatusOK

	json.NewEncoder(w).Encode(response)
}
