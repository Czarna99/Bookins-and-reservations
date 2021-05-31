package main

import (
	"Bookins-and-reservations/WebPage/pkg/handlers"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	fmt.Println("Application is running...")
	_ = http.ListenAndServe(portNumber, nil)
}
