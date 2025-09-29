package commands

import "os"

func MakeDir(name string) error {
	error := os.Mkdir(name, 0755);
	if error != nil {
		return error
	}
	return nil
}