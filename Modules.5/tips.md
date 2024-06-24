# Tips with packages and modules

https://go.dev/wiki/Modules#modules

- The first non-empty and uncommented line of any Go source code file contains the package specifier

- You can place several modules in the same repository, but you should not
  this is not a good idea. The content of each module is versioned as a single unit. Therefore, keeping two modules in the same repository would be
  would mean keeping separate versions for two different projects in the same repository.
  different projects

- When importing packages,
  that are not part of the standard library, you must specify the import path.
  import path. To get the import path, you combine the module path
  and the package path inside the module

- Instead of a name like `util`, it is better to give the package a name that describes the functionality it provides.
  Instead of a name like util, it is better to give the package a name that describes the functionality it provides. For example,
  suppose you have two helper functions, one to extract all the names from a string, and one to convert them to the required format.
  to extract all names from a string, and the other to convert them to the required format. In this case, you should not create two functions in the util package with the names ExtractNames and ExtractNames.
  package with the names ExtractNames and FormatNames. If you do this, the util names will be used at each
  call of these functions will use the names util.ExtractNames and util.FormatNames, where the prefix util does not give any information about what the functions do.
  functions

- we can add our specific name to imported modules

```go
import (
 crand "crypto/rand"
 "encoding/binary"
 "fmt"
 "math/rand"
)

```

- You can also use a dot (.) and underscore (\_) as a package name. When you use a dot (.), all exported identifiers of the imported package are moved into the namespace of the current package, so you can refer to them without using the
  prefix. This is not recommended, because it makes the source code less understandable and you will not be able to immediately understand where a particular identifier was defined.
  this or that identifier was defined: in the current or the imported
  package

# What's the difference between go run and go build

go run and go build are two commands used in the Go toolkit to compile and run Go programs.

1. go run: This command compiles and immediately runs Go source code, without saving the executable file on disk. This is useful for quickly running a program during development and debugging.

Example of using go run:

`go run main.go
`

2. go build: This command compiles the Go source code and saves an executable (binary file) in the current directory or at a specified location. This executable can be run directly without having to compile the source code each time.

An example of using go build:

```
go build -o myprogram
./myprogram
```

Thus, the main difference between go run and go build is that go run compiles and runs the program on the fly without saving the executable, while go build creates an executable that can be run at any time.

# How you should approach the organization of the module code

- https://www.youtube.com/watch?v=oL6JBUk6tj0

# Cyclical dependencies

- Go does not allow cyclic dependencies between packages. This means that if
  package A directly or indirectly imports package B, then package B cannot import package A directly or indirectly.
  directly or indirectly import package A:

If you have a cyclical addiction, you can choose one of
several options for dealing with this problem. In some cases, the cause
of this problem is that the code is divided too finely into packages. If two
packages depend on each other, it is likely that they should be combined
into a single package. In this case, we can solve the problem by combining the person and pet packages into one package.
the person and pet packages into a single package.
If you have a good reason to use two separate packages, it may be worth moving the person and pet packages out of each other.
packages, you might want to move only the items that give rise to the person and pet packages from one package to the other or to a new package.
package only those elements that create a cyclic dependency
