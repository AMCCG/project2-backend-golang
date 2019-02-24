package router

import (
	"net/http"
	"os"

	"github.com/AMCCG/project-2-backend-golang/router/api"
)

var PORT = os.Getenv("GOPORT")
var VERSION = "v1"
var CONTEXTPATH = "/api/" + VERSION

func Initial() {
	http.HandleFunc(CONTEXTPATH+"/users", api.UserHandler)
	http.ListenAndServe(":"+PORT, nil)
}
