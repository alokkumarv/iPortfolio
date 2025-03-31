package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func LoadTemplate() (*template.Template, error) {
	rootDir := "./template" // Change this to the directory you want to scan
	data := make([]string, 0)

	// Walk through the directory recursively
	err := filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Print file path
		if !d.IsDir() {
			fmt.Println(path)
			data = append(data, path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, path := range data {
		log.Printf("%v \n", path)
	}
	templ, err := template.ParseFiles(data...)
	return templ, err

}
