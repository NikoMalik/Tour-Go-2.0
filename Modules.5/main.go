package main

import (
	// we can import modules like this one
	"examplemodules/types" // gomodpackage/package
	"examplemodules/utilName"
	"fmt"
)

/********
### Program ###
*********/

// to runs program on go wee need
// => main() - entry function of the program
// always for main  -> ###package main###

// basic info how declare package to import

// package/library/module

// now check the utils.go file ->>>> utils.go

func main() {
	number := tint(52) // we can import function with u letter only with same packages
	// to import your value use Uppercase

	fmt.Println(number)
	fmt.Println("Hello Modules.5")
	name := utilName.GetRealName("Nikola") // we can import function with Uppercase

	user := types.User{ // we use package.function because its have another package
		Name: name,
		Age:  number,
	}
	fmt.Println(user)

}

// we can run this if we write go run main.go util.go

// or go build -o <package>
// and then ./<package>

// very important -> every file in root directory must have package main
/*
Go can only run single files not multiple files before initializing go mod
and can't build
*/
// go mod get all files and put them to binary

// but save all files in root directory is bad idea
