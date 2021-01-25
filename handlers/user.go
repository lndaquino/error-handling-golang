package handlers

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
)

func HandleUser(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Endpoint hit")

	_, err := parseBody(r)
	if err != nil {
		return AnotherErr(err)
	}

	_, err = getSomething(r)
	if err != nil {
		return BadRequestWithBody(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	return nil
}

func getSomething(r *http.Request) (string, error) {
	if (rand.Intn(100) % 2) == 0 {
		return "", errors.New("Invalid parameter")
	}
	return "something", nil
}

func parseBody(r *http.Request) (string, error) {
	if (rand.Intn(100) % 2) == 0 {
		return "", errors.New("Invalid body")
	}
	return "parsedBody", nil
}
