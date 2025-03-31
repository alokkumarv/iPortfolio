package routehandler

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, req *http.Request) {
	log.Printf("Triggered")
	temp, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	temp.Execute(w, nil)
}
