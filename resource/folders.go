package resource

import "os"

/** Create the Folders  */
func CreateFolders(args map[string]interface{}) {

	folderPath := args["folderPath"].(string)
	name := args["name"].(string)

	// Create the main Folder
	os.Mkdir(folderPath+name, 0777)

	// Create the dto folder and files
	os.Mkdir(folderPath+name+"/dtos", 0777)

	// Create the Entity Folder and File
	os.Mkdir(folderPath+name+"/entities", 0777)
}
