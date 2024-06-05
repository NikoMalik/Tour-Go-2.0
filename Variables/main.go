// Variables
package main

import "fmt" // we use this to printing in our program

func main() {
	eleven()
}

// Global variables always declared outside of the functions

// Examples with global variables

var Email string = "Q8QpH@example.com"

func nine() {
	fmt.Println(Email)
}

// we can also declare global variables without specifying the type

var Name = "Zhash"

func ten() {
	fmt.Println(Name)
}

// there is no difference in how variables are labeled and there
// will be no difference in performance
// it's up to you whether to use explicit typing or not.

// *if we want export our variable to other packages we have to start with a capital letter
// *we discuss this later

// if we want initial value with no value

var Email1 string //     but with type string we can use this variable in every our function and assign value there
var Number2 int   // 0 but with type int

func eleven() {
	fmt.Println(Email1)
	// output
	Email1 = "Q8QpH@example.com"
	fmt.Println(Email1)
	// output Q8QpH@example.com
	// we can declared variable with no value but with type string and assign value later in functions
	Email1 = 1
	fmt.Println(Email1)
	// syntax error because we can't assign another type to Email1
}

// golang always initialize variable with zero value

// We can also group variables in same block

var (
	Email2 string = "Q8QpH@example.com"
	Name1         = "Zhash"
	Name2         = "Zhash2"
)

func twelve() {
	fmt.Println(Email2)
	fmt.Println(Name1)
}

// use whatever you want

// we can also declare several variables
// of the same type:

var x, y int = 1, 2

//several variables of the same type with null values:

var z, w int

// or several variables of different types:

var a, b, c = 1, "hello", true

func severalv() {
	fmt.Println(a, b, c)
	// output 1 hello true
	fmt.Println(b, c)
	// output hello true
}

// Examples with local variables

// in local variables we don't need to specify the type and var is optional

// we can specify values with := only in local variable

func thirteen() {
	Email3 := "Q8QpH@example.com" // automatically infer the type string
	Number := 1                   // automatically infer the type int
	fmt.Println(Email3)
	fmt.Println(Number)
	// Email4 = "Q8QpH@example.com"

	// we can't use only = because is used to assign a value
}

// Constants examples

const PI = 3.14 // constant can't be changed and can't be nil or 0
// const string

// we can also declare type of constant

const Pi float64 = 3.14

// like in var we can group constants in same block

const (
	Email4  = "Q8QpH@example.com"
	Number3 = 1
)

// constants basically declared in top of your package

const yourself string = "nane"

func constants() {
	fmt.Println(yourself)
}
