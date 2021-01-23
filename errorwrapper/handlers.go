package errorwrapper

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponder interface {
	RespondError(w http.ResponseWriter, r *http.Request) bool
}

type BadRequestError struct {
	err  error
	body interface{}
}

func BadRequest(err error) *BadRequestError {
	return &BadRequestError{err: err}
}

func BadRequestWithBody(body interface{}) *BadRequestError {
	return &BadRequestError{body: body}
}

func (e *BadRequestError) RespondError(w http.ResponseWriter, r *http.Request) bool {
	if e.body == nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(e.body); err != nil {
			log.Printf("Failed to encode a response: %v", err)
		}
	}

	return true
}

func (e *BadRequestError) Error() string {
	return e.err.Error()
}
