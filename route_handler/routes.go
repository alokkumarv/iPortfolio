package routehandler

import (
	"log"
	"net/http"

	"github.com/alokkumarv/utils"
)

func Home(w http.ResponseWriter, req *http.Request) {
	log.Printf("Triggered")
	templ, err := utils.LoadTemplate()
	if err != nil {
		log.Printf("Error occured while parsing files %v", err)
	}
	log.Printf("%v", templ.DefinedTemplates())

	err = templ.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Printf("Error occured while executing page %v", err)
	}
}
