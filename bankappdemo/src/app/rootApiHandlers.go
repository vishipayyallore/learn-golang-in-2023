package app

import (
	"fmt"
	"net/http"
)

func apiRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Bank App Web API in Go!")
}
