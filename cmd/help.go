package cmd

import "fmt"

/** Print in console the Help Options*/
func Help() {
	fmt.Println(`Command List:
	All the flags: -af | -all-flags;
	Flag to add the Authentication: -a | -auth;
	Flag to add the Documentation: -d  | -docs;
	Flag to add the Logger: -l | -log
	
	Avaible Types: number | string | email | password
	Optional Declaration "?:"
	Required Declaration ":"
	The directory path should be the last argument to the command
	`)
}
