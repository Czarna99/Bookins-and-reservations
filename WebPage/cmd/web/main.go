package main

import (
	"Bookins-and-reservations/WebPage/pkg/config"
	"Bookins-and-reservations/WebPage/pkg/handlers"
	"Bookins-and-reservations/WebPage/pkg/render"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateChace = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	fmt.Println("Application is running...")

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
