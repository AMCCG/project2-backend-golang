package constant

import "os"

var PORT = os.Getenv("GOPORT")
var VERSION = "v1"
var CONTEXTPATH = "/api/" + VERSION

// CONSTANT URL
var UsersURL = CONTEXTPATH + "/users/"
