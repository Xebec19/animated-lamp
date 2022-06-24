package main

import (
	"animated-lamp/pkg/config"
	"animated-lamp/pkg/handlers"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	repo := handlers.NewRepo(&app)
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// http.HandleFunc("/temp", handlers.Repo.Divide)
	fmt.Printf(fmt.Sprintf("Starting application on port %s", portNumber))
	// _ = http.ListenAndServe(portNumber, nil)
	src := &http.Server{
		Addr:     portNumber,
		Handlers: routes(app),
	}
	err = src.ListenAndServe()
	log.Fatal(err)
}
