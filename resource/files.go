package resource

import (
	"fmt"
	"os"

	"github.com/yc4rlos/go-nest-files/go-nest-files/files"
)

/*Create all files */
func CreateFiles(titledName, name, folderPath string, auth, documentation, logger bool, properties []string) {

	controller := files.ControllerFile(titledName, name, auth, documentation, logger)
	service := files.ServiceFile(titledName, name, auth, documentation, logger)
	module := files.ModuleFile(titledName, name)
	createDto := files.CreateDtoFile(titledName, name, auth, documentation, logger, properties)
	updateDto := files.UpdateDtoFile(titledName, name)
	entity := files.EntityFile(titledName, name, properties)

	create(folderPath, name, controller, service, module, createDto, updateDto, entity)

	if documentation {
		dto := files.DtoFile(titledName, name)
		create(folderPath, name, dto)
	}
}

/** Create the file in the folders */
func create(folderPath, name string, list ...[2]string) {

	for _, props := range list {
		name := props[0]
		content := props[1]

		file, err := os.Create(folderPath + name)
		if err != nil {
			fmt.Println(err)
			return
		}

		fileBytes := []byte(content)
		file.Write(fileBytes)

		file.Close()
	}

}
