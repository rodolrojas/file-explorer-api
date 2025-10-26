package commands

import (
	"api/helpers"
	"os"
)

func CreateFile(path string, filename string) error {
	fullPath := helpers.BuildPath(path, filename)
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}