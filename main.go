package main

import (
	"log"

	"github.com/AMCCG/project-2-backend-golang/router"
)

func main() {
	log.Print("************ Starting Service ************")
	log.Print("PORT : " + router.PORT)
	log.Print("CONTEXTPATH : " + router.CONTEXTPATH)
	router.Initial()
}
