package endpoints

import (
	"net/http"
	"fmt"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index")
	w.WriteHeader(http.StatusOK)
}