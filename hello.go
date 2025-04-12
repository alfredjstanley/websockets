package hello

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func RunServer() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}