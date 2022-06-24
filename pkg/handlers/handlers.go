package handlers

import (
	"animated-lamp/pkg/config"
	"animated-lamp/pkg/render"
	"errors"
	"fmt"
	"net/http"
)

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

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
