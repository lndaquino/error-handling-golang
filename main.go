package main

import (
	"error-wrapper/handlers"
	"fmt"
	"net/http"
)

func main() {
	setupRoutes()

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func setupRoutes() {
	http.HandleFunc("/user", handlers.WithError(handlers.HandleUser))
	http.HandleFunc("/get", testGet)
}

func testGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
