package utils

import (
	"fmt"
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
			data = append(data, path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
	templ, err := template.ParseFiles(data...)
	return templ, err

}
