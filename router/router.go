package router

import (
	"net/http"

	"github.com/AMCCG/project-2-backend-golang/constant"
	"github.com/AMCCG/project-2-backend-golang/router/api"
)

func Initial() {
	http.HandleFunc(constant.UsersURL, api.UserHandler)
	http.ListenAndServe(":"+constant.PORT, nil)
}
