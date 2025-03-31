package routehandler

import (
	"log"
	"net/http"

	"github.com/alokkumarv/repository"
	"github.com/alokkumarv/utils"
)

func Home(w http.ResponseWriter, req *http.Request) {
	log.Printf("Triggered")
	templ, err := utils.LoadTemplate()
	if err != nil {
		log.Printf("Error occured while parsing files %v", err)
	}
	user := repository.GetUserInfo()
	if user == nil {
		log.Printf("Not able to get user data ")
	}
	log.Printf("user : %v", user)
	err = templ.ExecuteTemplate(w, "index.html", user)
	if err != nil {
		log.Printf("Error occured while executing page %v", err)
	}
}
