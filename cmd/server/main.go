package main

import (
	"election-voting-system/internal/routes"
	"fmt"
	"log"
	"net/http"
)
func main() {
	routes.RegisterRoutes()

	fmt.Println("Server running on http://localhost:8080/")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
