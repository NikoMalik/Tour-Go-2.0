// inside new folder we create new file user.go
// what would be the package name?
// create name package as same as folder to be clear

package types

type User struct {
	Name string // we need declare this with uppercase to use it in other packages
	Age  byte
}

// name - only for same packages -> private access
// Name -  available in all packages -> public access
