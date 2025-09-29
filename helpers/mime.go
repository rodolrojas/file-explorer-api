package helpers

import (
	"io/fs"
	"mime"
	"path/filepath"
)

func GetMimeType(file fs.FileInfo) string {
	// Implementation for getting MIME type based on file extension
	ext := filepath.Ext(file.Name())
	mimeType := mime.TypeByExtension(ext)
	if mimeType != "" {
		return mimeType
	}
	return ""
}