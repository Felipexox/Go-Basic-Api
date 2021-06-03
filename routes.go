package main

import (
	"fmt"
	"net/http"
)

func handleRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Batata")
}
