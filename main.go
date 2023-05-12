package main

import (
	"fmt"

	"github.com/yc4rlos/go-nest-files/cmd"
	"github.com/yc4rlos/go-nest-files/resource"
)

func main() {

	args := map[string]interface{}{}

	cmd.GetArgs(args)

	resource.CreateFolders(args)
	resource.CreateFiles(args)

	fmt.Println("Files generated!")
}
