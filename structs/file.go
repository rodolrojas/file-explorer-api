package structs

import (
	"time"
)

type ApiFileStruct struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Size int64 `json:"size"`
	Type string `json:"type"`
	IsDir bool `json:"is_dir"`
	ModTime time.Time `json:"mod_time"`
	User string `json:"user"`
	Group string `json:"group"`
	Mode string `json:"mode"`
}