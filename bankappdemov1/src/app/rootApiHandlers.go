package app

import (
	"fmt"
	"net/http"
)

func ApiRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Bank App Web API in Go!")
}
