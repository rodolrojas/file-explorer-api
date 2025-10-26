package commands

import (
	"api/helpers"
	"os"
)

func MakeDir(path string, name string) error {
	fullPath := helpers.BuildPath(path, name)
	error := os.Mkdir(fullPath, 0755)
	if error != nil {
		return error
	}
	return nil
}