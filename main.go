package main

import (
	"api/src/config.go"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.Load()

	r := router.GenerateRouter()

	fmt.Println("Running API...")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))

}
