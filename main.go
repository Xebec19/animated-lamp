package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/temp", Divide)
	_ = http.ListenAndServe(":8080", nil)
}
