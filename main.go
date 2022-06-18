package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum, err := addValues(2, 2)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2+2 is %d", sum))
	}
}

// addValues add two integers
func addValues(x, y int) (int, error) {
	var sum int
	sum = x + y
	return sum, nil
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	_ = http.ListenAndServe(":8080", nil)
}
