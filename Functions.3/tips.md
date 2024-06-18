# Tips and Good Practices

https://go.dev/blog/ismmkeynote

- When a function accepts several parameters of the same type, their type
  can be specified as follows:
  `func div(numerator, denominator int) int `

- !When returning multiple values, multiple values are always returned
  multiple values

- Use the `_` identifier whenever you do not need the value returned by the function

- If your function returns values, never use an empty
  return statement. This can make it very difficult to understand
  exactly what value it returns

- Passing functions as parameters to other functions is often used to perform different operations on data of the same
  types

- Never give a variable or
  function name nil

- Noticing that the function returns
  a pointer to a local variable, it places the value of this variable in the slices

- You can embed any type into a structure, not just another structure.
  This allows you to elevate the methods of the embedded type to the structure containing it
  structure

- Since you can assign literal expressions to constants, you may
  may encounter code samples that suggest using iota
  as follows:
  `type BitField int
const (
 Field1 BitField = 1 << iota // assign 1
 Field2 // assigned 2
 Field3 // assigned 4
 Field4 // assigned 8
)`
  No matter how clever and advanced this solution may seem, be extremely careful and document your actions.
  be very careful and document your actions if you decide to use this pattern.

- In principle, any simple operator can be placed in front of an if statement condition, including things like calling a non-returning function or assigning a new value to an existing variable.
  operator, including such things as calling a function that does not return a value or assigning a new value to an existing variable. However
  this is not a good idea. To avoid confusion, use this feature only to define new variables whose scope will be limited to if/else statements.
  Also, do not forget that, as in any other block, a variable,
  declared inside an if statement will shadow variables of the same name declared in other blocks.
  variables with the same name declared in other blocks

- The for-range loop can only be used to bypass built-in composite types or custom types based on them.
  types or custom types based on them

- The for-range loop copies the values of the elements

- Prefer a for-range loop when you need to loop through the contents of an instance of one of the built-in composite types. This allows you to avoid cumbersome template code that you need to write when traversing arrays, slices and maps with other types.
  when traversing arrays, slices, and maps with other types of for loops.
  for loop.

- Prefer a for loop best used
  when it is not necessary to traverse the entire contents of a composite type
  first to last element

- Although you can also specify a function with some return values in the defer statement, there is no way you can read those values.
  `func example() {
 defer func() int {
 return 2 // this value cannot be read
 }
}`

- Beginning Go developers often forget to put parentheses,
  when specifying a closure in the defer statement. Since the absence of these
  parentheses causes a compile-time error, you'll gradually get used to it.
  to the correct version. However, it is useful to remember that parentheses are necessary to keep the
  parentheses are necessary to indicate the values passed to the closure
  when the closure is started

# Difference in performance difference between `for-range` and `for-loop`

#### As you know `for-range` copy every element and if we make test in performance we get this

```go
func sumRange(objects []Obj) int {
  ret := 0
  for _, v := range objects {
    ret += v.index
  }
  return ret
}

func sumLoop(objects []Obj) int {
  ret := 0
  for i := 0; i < len(objects); i++ {
    ret += objects[i]
  }
  return ret
}
```

Result :

BenchmarkForRange-4 443161 2371ns/op
BenchmarkForLoop-4 1863501 641.7ns/op

> the bigger the object inside the slide, the bigger the difference will be.
