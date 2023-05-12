package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/yc4rlos/go-nest-files/go-nest-files/cmd"
	"github.com/yc4rlos/go-nest-files/go-nest-files/resource"
)

/** Defines if it will have authentication*/
var auth bool

/** Defines if it will have documentation*/
var documentation bool

/** Defines if have error log config (With Nest Winston)*/
var logger bool

/** Resource Name*/
var name string

/** Titled Resource Name*/
var titledName string

/** Properties List*/
var properties = make([]string, 2, 15)

/** Folder Path */
var folderPath string

func main() {

	getArgs()

	resource.CreateFolders(folderPath, name)
	resource.CreateFiles(titledName, name, folderPath, auth, documentation, logger, properties)

	fmt.Println("Files generated!")
}

func getArgs() {
	for i, arg := range os.Args {
		if i == 0 {
			continue
		}

		if arg[0] == '-' {
			switch arg {
			case "-af", "-all-flags":
				auth = true
				documentation = true
				logger = true
			case "-a", "-auth":
				auth = true
			case "-d", "-docs":
				documentation = true
			case "-l", "-log":
				logger = true
			case "-h", "-help":
				cmd.Help()
				return
			default:
				fmt.Printf("Command %s invalid, use -h for more information.\n", arg)
			}
		} else {
			if name == "" {
				name = arg
				titledName = strings.Title(arg)
				continue
			}

			if i == len(os.Args)-1 {

				if strings.ContainsRune(arg, '/') || strings.ContainsRune(arg, '\\') {
					folderPath = arg
				} else {
					fmt.Println("The Path cannot be empty.")
					os.Exit(0)
				}
			} else if strings.ContainsRune(arg, ':') {
				properties = append(properties, arg)
			}
		}
	}
}
