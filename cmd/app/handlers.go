package app

import (
	"fmt"
	"net/http"
)

func rootHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, http.StatusOK)
	})
}
