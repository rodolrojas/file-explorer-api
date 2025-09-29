package responses

import "api/structs"

type ListFilesReturn struct {
	Items []structs.ApiFileStruct `json:"items"`
	Count int				   `json:"count"`
}