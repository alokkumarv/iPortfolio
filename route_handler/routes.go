package routehandler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/alokkumarv/models"
)

func Home(w http.ResponseWriter, req *http.Request) {
	log.Printf("Triggered")
	temp, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
	}
	user := models.User{Name: "Alok Kumar Verma"}
	log.Printf("Data %v", user.Name)
	err = temp.Execute(w, user)
	if err != nil {
		log.Printf("Error occured while executing page %v", err)
	}
}
