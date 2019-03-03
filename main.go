package main

import (
	"log"

	"github.com/AMCCG/project-2-backend-golang/router"

	"github.com/AMCCG/project-2-backend-golang/constant"
)

func main() {
	log.Print("************ Starting Service ************")
	log.Print("PORT : " + constant.PORT)
	log.Print("CONTEXTPATH : " + constant.CONTEXTPATH)
	router.Initial()
}
