package structs

import (
	"time"
)

type ApiFileStruct struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Size int64 `json:"size"`
	Type string `json:"type"`
	IsDir bool `json:"isDir"`
	ModTime time.Time `json:"modTime"`
	User string `json:"user"`
	Group string `json:"group"`
	Mode string `json:"mode"`
}