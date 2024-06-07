// Types
// https://golang.org/ref/spec#Numeric_types
// https://www.asciitable.com/
package main

import "fmt"

var (
	ToBe         bool    = false // default value is false for bool
	floatdefault float64         // 0 by default
	floatVar     float32 = 3.14  // 4 bytes (32 bit)   3,40282346638528859811704183484516925440e+38 to 1,401298464324817070923729583289916131280e–45
	floatVar2    float64 = 3.14  // 8 bytes (64 bit)  1,797693134862315708145274237317043567981e+308 to 4,940656458412465441765687928682213723651e-324

	name       string = "walter"
	defaultage int         // 0 by default
	age        int    = 25 // if you are have 64 bit system it will be 64 bit or if you have 32 bit system it will be 32 bit
	age4       int8   = 25 // -128 to 127 1 byte (8 bit)
	age5       int16  = 25 // -32768 to 32767 2 bytes (16 bit)
	age2       int32  = 25 // -2147483648 to 2147483647 4 bytes (32 bit)
	age3       int64  = 25 // -9223372036854775808 to 9223372036854775807 8 bytes (64 bit)

	uintage  uint   = 25 // same like int but with only positive numbers (0 to _64/_32-bit)
	uintage4 uint8  = 25 // same as byte (0 to 255) 1 byte (8 bit)
	uintage5 uint16 = 25 // 0 to 65535 2 bytes (16 bit)
	uintage2 uint32 = 25 // 0 to 4294967295 4 bytes (32 bit)
	uintage3 uint64 = 25 // 0 to 18446744073709551615 8 bytes (64 bit)
	a        byte   = 62 // 0 to 255 byte is alias for uint8 1 byte (8 bit)
	// you will rarely see the name uint8 in Go-code:
	// instead, it is common to byte

	runeVar rune = 'a' // alias for int32 4 bytes (32 bit)
	// "" = string '' = symbol unicode codepoints
	// fmt.Println(runeVar) = 97 // output bit https://www.asciitable.com/

// The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific
// reason to use a sized or unsigned integer type.

)

/********
###  Struct ###
*********/

// struct to make an analogy its like a class in python
// class User:
// 	def __init__(self, name, id, age):
// 		self.name = name
// 		self.id = id
// 		self.age = age

// but this is not a class and not oop
// struct can store any fields with any types and he may have different methods
// and create objects based on this structure
var User = struct {
	name string
	id   uint
	age  byte // age can't be negative and can't be more than 255 we can use byte for this
}{"walter", 1, 25}

func userik() {
	fmt.Println(User)
} // output {walter 1 25}

// but we can't reuse it and its rare

// usually we use struct  with type 99% cases

/********
### Type  ###
*********/

type UserType struct {
	name string
	id   uint
	age  byte
}

// we can use something  like this
func PrintUserType() {
	user := UserType{
		"walter",
		1,
		25,
	}
	// or we can set values with this

	user = UserType{
		name: "walter",
		id:   1,
		age:  25,
	}

	fmt.Println(user)

	// output {walter 1 25}

	//lets use the printf to print with another format

	fmt.Printf("%+v\n", user) // https://pkg.go.dev/fmt you can check all about fmt and values here
	//and we can access fields directly

	fmt.Println(user.name)
	fmt.Println(user.id)
	fmt.Println(user.age)

	// output
	// walter
	// 1
	// 25

}

// type can be used to create type of object
// type creates a data type, not a variable
type MyType int // alias for int

var tip MyType = 1 // we create a variable of type MyType and assign it to 1
// another example

type Windows string
type Host string

var hosts = map[Windows]Host{
	"win1": "host1",
	"win2": "host2",
	"win3": "host3",
}

/********
### Maps ###
*********/

// Maps like dictionaries
// map[Type|key]Type|value key-value pairs

var myMap = map[string]int{ // map[Type|key]Type|value
	"key1": 1,
	"key2": 2,
}

//key is a unique value and cannot be duplicated

// we can also save the empty map
// initialize an empty map
var myMap2 = map[int]string{} // empty map == nil

// we can append values to the map

func PrintMap() {
	myMap2[1] = "key1"
	myMap2[2] = "key2"
	myMap2[3] = "key3"
	fmt.Println(myMap2)

	// we use key to access to maps values
	fmt.Println(myMap2[1])
	// output key1

	// if we try to refer to nonexistent values, we will get  output == "" if value is string
	fmt.Println(myMap2[0])

	// we can check whether the current item exists in the map
	key, ok := myMap2[1] // we use _ to ignore the value
	if !ok {             // ok is boolean
		fmt.Println("key1 does not exist")
	}

	fmt.Println(key) // key1

	nonexistentKey, ok := myMap2[0]
	if !ok { // ok is boolean
		fmt.Println("key0 does not exist")
	}

	fmt.Println(nonexistentKey) // key0 does not exist

	// we can delete items from the map
	delete(myMap2, 1) // delete(map, key)
	fmt.Println(myMap2)

	// output map[2:key2 3:key3]

	myMap2[4] = "key4"
	// we can also iterate on maps

	for key, value := range myMap2 { //
		fmt.Printf("%d %s\n", key, value)
	}
	// output 3 key3 4 key4 2 key2
	// we can use range to iterate on maps with only keys
	for key := range myMap2 {
		fmt.Println(key) // ouput 2 3 4

	}
}

// we can make a map like with another way

var myMap3 = make(map[string]int) // make
//Maps created with the make function still have zero length
// and are not limited in their growth by the originally specified size

// we can also declare length of map
var myMap4 = make(map[string]int, 10) // make(map[Type|key]Type|value, length|10)

// example using maps
var users = map[string]UserType{ // we use map[Type|key]Type|value value = UserType { name: string, id: uint, age: byte }
	"walter": {
		"walter",
		1,
		25,
	},
	"walter2": {
		"walter2",
		2,
		26,
	},
}

/********
### Arrays ###
*********/

// *we rarely use arrays in go

// all elements in massive must be of the same type
//we can declare arrays in different ways
// you should always specify the number of values in the arrays
var x [3]int // 1 way

// x = (x[0], x[1], x[2]) // 2 way
// in the array, the counting always starts from the zero value
// 2 way
var x2 = [5]int{1, 2, 3, 4, 5} // we use [quantity]Type|int{value|1,2,3,4,5}
// 3 way
// we can  specify only the indices of individual elements with corresponding values
var x3 = [12]int{1, 3: 2, 5: 4} // we use [quantity]Type|int{value|1,0,0,0,0,2,0,0,0,0,0,4}

// 4 way

var x4 = [...]int{1, 2, 3, 4, 5} // we use [...]Type|int{value|1,2,3,4,5}
// ... = golang during compilation gives the amount of elements in the array
// out [5]int{1, 2, 3, 4, 5}
// You're omitting the size of the array and letting the compiler figure it out for you
func comparative() {
	y := [5]int{1, 2, 3, 4, 5}
	fmt.Println(x4 == y)

	// output true
}

// we can also declare multidimensional arrays
var x5 = [3][2]int{{1, 2}, {3, 4}, {5, 6}} // we use [quantity][quantity]Type|int{{value|1,2},{value|3,4},{value|5,6}}

// we can read length of arrays
func PrintArray() {
	fmt.Println(len(x4))
	// output 5
}

// we can add values to arrays with this
func addarray() {
	x := [2]int{}

	x[0] = 1

	x[1] = 2
	fmt.Println(x)
	// output [1 2]
}

// in Go, the size
// of an array is considered to be a part of its type

/* This also means that you cannot use a variable
to specify the size of an array, because types must be defined at compile time, not at runtime.
at compile time, not at runtime
*/
// another way to declare an comments /* */

/********
### Slices ###
*********/

// most popular data structure to accommodate a value sequence

var slices = []string{"a", "b", "c"} // we use []Type|string{"value|a","value|b","value|c"}
// slices is a reference to an array
//slices is dynamic and can grow or shrink

// we can declare slices in different ways
var slices1 = []string{}              // 1 way most popular way  to initialize empty slices
var slices2 = []string{"a", "b", "c"} // 2 way
var slices3 = make([]string, 3)       // 3 way = make([]Type|value, length|3)
var slices4 = make([]string, 3, 5)    // 4 way = make([]Type|value, length|3, capacity|5)
// capacity is .....

/*The capacity of a slice is the number of elements that the slice can hold. It is not the same as the length of the slice, which is the number of elements that are currently in the slice.

Slices are created with a capacity that is equal to the length, but the capacity can be increased if more elements are added to the slice. The capacity is increased automatically as needed to accommodate new elements.

*/

// The capacity of a slice can be retrieved using the `cap` function. For example:

func checkcapacity() {
	s := make([]int, 5)
	fmt.Println(cap(s)) // prints 5

	// The capacity of a slice can be increased using the `append` function. For example:

	s = append(s, 6)
	fmt.Println(cap(s)) // prints 10

	//append is is a function in
	//Go that is used to add one or more elements to a slice
	// It returns a new slice with the added elements

	s = append(s, 7, 8, 9)
	/* One important thing to note is that `append`
	does not modify the original slice.
	Instead, it returns a new slice with the added elements.
	This means that if you want to modify the original slice
	you need to assign the result of `append` to the original slice
	*/
	fmt.Println(cap(s)) // prints 10

	s = append(s, 10, 11, 12)
	fmt.Println(cap(s)) // prints 20 because
	/*the capacity of a slice is usually doubled each time
	new elements are added to it.
	This is because Go uses the array mechanism to implement
	slices, and increasing capacity in this way is an efficient way
	to accommodate new items.
	However, there are some exceptions to this rule.
	For example, if the capacity of a slice is already very large
	(e.g., greater than 1024), the capacity may not double when new
	elements are added. Also, if the slice was created using the make function
	with a specified capacitance value, the capacitance will not double
	even if the capacitance is exceeded.


	*/
	fmt.Println(s)
	// output [0 0 0 0 0 6 7 8 9 10 11 12]
	// we have 5 zero values because we initialized it with 5 empty values
}

func main() {
	// PrintMap()
	// PrintArray()
	// addarray()
	checkcapacity()
}
