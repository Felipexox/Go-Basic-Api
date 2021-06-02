package hq

import (
	"fmt"
	"net/http"
)

func handleRoutes() {
	http.HandleFunc("/allHQ", getAllHQ)
}

func getAllHQ(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get All HQ")
}
