package main

import (
	"log"
	"project-2-backend-golang/constant"
	"project-2-backend-golang/router"
)

func main() {
	log.Print("************ Starting Service ************")
	log.Print("PORT : " + constant.PORT)
	log.Print("CONTEXTPATH : " + constant.CONTEXTPATH)
	router.Initial()
}
