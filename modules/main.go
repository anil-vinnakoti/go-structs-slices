package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	str := color.CyanString("hello world!")
	fmt.Println(str)	
}

// to initiate a new module 						=> go mod init "repo_path_and_name_of_module"
// to add new modules to repo 						=> go get "module_path"
// to clean up unused and add using modules 		=> go mod tidy
// list all modules(including go in-build) 			=> go list all
// list all modules(dependency and downloaded) 		=> go list -m all