package repository

import (
	"encoding/json"
	"log"
	"os"

	"github.com/alokkumarv/models"
)

func GetUserInfo() *models.User {
	var user models.User
	file, err := os.Open("metadata.json")
	if err != nil {
		log.Printf("Error reading Config file %v", err)
		return nil
	}
	err = json.NewDecoder(file).Decode(&user)
	if err != nil {
		log.Printf("Error while decoding metadata file")
		return nil
	}
	return &user
}
