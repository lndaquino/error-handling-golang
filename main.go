package main

import (
	"error-wrapper/errorwrapper"
	"errors"
	"fmt"
	"net/http"
)

func main() {
	setupRoutes()

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func setupRoutes() {
	http.HandleFunc("/user", errorwrapper.WithError(handleThing))
	http.HandleFunc("/get", testGet)

}

func handleThing(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Endpoint hit")

	_, err := getSomething(r)
	if err != nil {
		return errorwrapper.BadRequestWithBody(err.Error())
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func getSomething(r *http.Request) (string, error) {
	return "", errors.New("Invalid parameter")
}

func testGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
