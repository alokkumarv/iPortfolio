package main

import (
	"log"
	"net/http"
	"os"

	routehandler "github.com/alokkumarv/route_handler"
)

func main() {
	// Server Static files
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/assets"))))

	http.HandleFunc("/", routehandler.Home)
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	log.Printf("Server is running at endpoint %v:%v", host, port)
	if err := http.ListenAndServe(host+":"+port, nil); err != nil {
		log.Printf("Error while running server ")
	} else {
		log.Printf("Server is running at endpoint %v:%v", host, port)
	}

}
