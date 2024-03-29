package main

import (
	"app-web/handlers"
	"log"
	"net/http"
)
	

func main() {

	mux :=  http.NewServeMux()
	// register handlers for different routes here
	mux.HandleFunc("/",handlers.HomeHandler)
	mux.HandleFunc("/home", handlers.HomeHandler) // same as "/"
	mux.HandleFunc("/about", handlers.AboutHandler)
	mux.HandleFunc( "/product", handlers.ProductHandler)
	mux.HandleFunc("/post-get", handlers.PostGet)
	mux.HandleFunc("/form", handlers.Form)
	mux.HandleFunc("/process", handlers.Process)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer)) 

	log.Println("Starting web on port 8080")

	err:= http.ListenAndServe( ":8080", mux )
	log.Fatal(err)
}
