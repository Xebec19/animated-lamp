package handlers

import (
	"animated-lamp/pkg/render"
	"errors"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum, err := addValues(2, 2)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2+2 is %d", sum))
	}
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", f))
}

// addValues add two integers
func addValues(x, y int) (int, error) {
	var sum int
	sum = x + y
	return sum, nil
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}
