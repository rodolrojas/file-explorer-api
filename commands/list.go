package commands

import (
	"os"

	"api/helpers"
	"api/responses"
	"api/structs"
	"log"
)

func ListFiles(path string, filterDirs bool) (responses.ListFilesReturn, error) {
	var fileList []structs.ApiFileStruct
	files, err := os.ReadDir(path)
	
	if err != nil {
		log.Println(err)
		log.Println("Error reading directory:", err)
		return responses.ListFilesReturn{}, err
	}

	for _, file := range files {
		if filterDirs && !file.IsDir() {
			continue
		}
		currentFileInfo, err := file.Info()
		if err != nil {
			log.Println("Error getting file info:", err)
			return responses.ListFilesReturn{}, err
		}
		currentFileUserName, currentFileUserGroup, err := helpers.GetUnixUserGroup(currentFileInfo)
		
		if err != nil {
			log.Println("Error getting file info:", err)
			return responses.ListFilesReturn{}, err
		}
		var currentFile = structs.ApiFileStruct{
			Name : file.Name(),
			Size : currentFileInfo.Size(),
			IsDir : currentFileInfo.IsDir(),
			Mode : currentFileInfo.Mode().String(),
			User : currentFileUserName,
			Group : currentFileUserGroup,
			ModTime: currentFileInfo.ModTime().UTC(),
			Path : helpers.BuildPath(path, file.Name()),
			Type : helpers.GetMimeType(currentFileInfo),
		}
		fileList = append(fileList, currentFile)
	}
	return responses.ListFilesReturn{
		Items: fileList,
		Count: len(fileList),
	}, nil
}	