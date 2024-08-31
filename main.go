package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Running API...")

	r := router.GenerateRouter()

	log.Fatal(http.ListenAndServe(":8080", r))

}
