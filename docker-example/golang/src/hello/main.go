package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		name := r.URL.Query().Get("name")

		fmt.Fprintf(w, "Hello, %s", name)

	}))
}
