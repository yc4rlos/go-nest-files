package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/gertd/go-pluralize"
)

func GetArgs(args map[string]interface{}) {
	args["properties"] = make([]string, 2, 15)

	if len(os.Args) <= 2 {
		fmt.Println("Is necessary provide a folder name and at last 1 properties\n=================")
		Help()
	}

	for i, arg := range os.Args {
		if i == 0 {
			continue
		}

		if arg[0] == '-' {
			switch arg {
			case "-af", "-all-flags":
				args["auth"] = true
				args["documentation"] = true
				args["logger"] = true
			case "-a", "-auth":
				args["auth"] = true
			case "-d", "-docs":
				args["documentation"] = true
			case "-l", "-log":
				args["logger"] = true
			case "-h", "-help":
				Help()
				os.Exit(0)
			default:
				fmt.Printf("Command %s invalid, use -h for more information.\n", arg)
			}
		} else {
			if args["name"] == "" {
				args["name"] = arg

				p := pluralize.NewClient()

				args["singularName"] = p.Singular(args["name"].(string))
				continue
			}

			if i == len(os.Args)-1 {

				if strings.ContainsRune(arg, '/') || strings.ContainsRune(arg, '\\') {
					args["folderPath"] = arg
				} else {
					fmt.Println("The Path cannot be empty.")
					os.Exit(0)
				}
			} else if strings.ContainsRune(arg, ':') {
				args["properties"] = append(args["properties"].([]string), arg)
			}
		}
	}
}
