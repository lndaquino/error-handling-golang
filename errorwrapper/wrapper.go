package errorwrapper

import (
	"log"
	"net/http"
)

type HandlerE = func(w http.ResponseWriter, r *http.Request) error

func WithError(h HandlerE) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			if er, ok := err.(ErrorResponder); ok {
				if er.RespondError(w, r) {
					return
				}
			}

			log.Printf("Something went wrong: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}
}
