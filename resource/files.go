package resource

import (
	"fmt"
	"os"

	"github.com/yc4rlos/go-nest-files/files"
)

/*Create all files */
func CreateFiles(args map[string]interface{}) {

	// Values
	name := args["name"].(string)
	singularName := args["singularName"].(string)
	folderPath := args["folderPath"].(string)

	auth := args["auth"].(bool)
	documentation := args["documentation"].(bool)
	logger := args["logger"].(bool)

	properties := args["properties"].([]string)

	// Files
	controller := files.ControllerFile(name, singularName, auth, documentation, logger)
	service := files.ServiceFile(name, singularName, auth, documentation, logger)
	module := files.ModuleFile(name, singularName)
	createDto := files.CreateDtoFile(name, singularName, auth, documentation, logger, properties)
	updateDto := files.UpdateDtoFile(name, singularName)
	entity := files.EntityFile(name, singularName, properties)

	create(folderPath, name, controller, service, module, createDto, updateDto, entity)

	if documentation {
		dto := files.DtoFile(name, singularName)
		create(folderPath, name, dto)
	}
}

/** Create the file with the provided name and value */
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
