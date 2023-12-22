package handlers

import (
	"fmt"
	"net/http"
	"os"
)

var f, _ = os.ReadFile("static/index.html")

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string(f))
}
