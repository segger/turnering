package main

import (
	"fmt"
	"log"
	"net/http"
	"server/router"
)

func main() {
	r := router.Router()

	fmt.Println("Starting server on 8084...")
	log.Fatal(http.ListenAndServe(":8084", r))
}
