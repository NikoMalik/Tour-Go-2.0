https://github.com/golang/go/blob/f296b7a6f045325a230f77e9bda1470b1270f817/src/runtime/runtime2.go#L395

go build -gcflags='-m=3' // for debug info

# There are three immutable rules for memory allocation from version to version of Golang. You have 100% allocated value on the hip if:

1. The result is returned by reference;

2. The value is passed to an argument of type interface{} - the fmt.Println argument;

3. The size of the variable value exceeds the stack limits.
