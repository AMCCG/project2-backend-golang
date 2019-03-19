package router

import (
	"net/http"

	"project-2-backend-golang/constant"
	"project-2-backend-golang/router/api"
)

func Initial() {
	http.HandleFunc(constant.UsersURL, api.UserHandler)
	http.ListenAndServe(":"+constant.PORT, nil)
}
