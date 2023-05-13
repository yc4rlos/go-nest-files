package main

import (
	"fmt"

	"github.com/yc4rlos/go-nest-files/cmd"
	"github.com/yc4rlos/go-nest-files/resource"
)

func main() {

	args := map[string]interface{}{
		"auth":          false,
		"documentation": false,
		"logger":        false,
		"name":          "",
		"singularName":  "",
		"folderPath":    "",
		"properties":    make([]string, 2, 15),
	}

	cmd.GetArgs(args)

	resource.CreateFolders(args)
	resource.CreateFiles(args)

	fmt.Println("Files generated!")
}
