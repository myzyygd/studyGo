package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	})

	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "World")
	})

	http.ListenAndServe(":8080", nil)
}
