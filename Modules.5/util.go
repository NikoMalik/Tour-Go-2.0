package main // need to be package main because declared in root directory
// import doesn't work also if we use same package we need declare go mod

// you can check what happens if we declare another package in root directory
func tint(i byte) byte {
	return i
}

// to import this function in any file in your go project

// we should use ```go mod init <name>```

// In Go, the path to the repository where the module is located is usually used as the identifier,
// where the module is located like github.com/user/repo

// name for module doesn't matter

// to check commands use go mod help

// and we have it

/*
go mod init modules                                                            [9:05:32]
go: creating new go.mod: module modules
go: to add module requirements and sums:
        go mod tidy

*/
