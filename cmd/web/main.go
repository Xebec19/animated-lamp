package main

import (
	"animated-lamp/pkg/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/temp", handlers.Divide)
	_ = http.ListenAndServe(":8080", nil)
}
