package helpers

import "strings"

func BuildPath(basePath, relativePath string) string {
	if basePath == "/" {
		return "/" + relativePath
	}
	basePath = strings.TrimRight(basePath, "/")
	return basePath + "/" + relativePath
}