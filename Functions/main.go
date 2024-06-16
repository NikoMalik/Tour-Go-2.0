package main

import (
	"context"
	"database/sql"
	"fmt"
	"sort"
)

/********
### Func ###
*********/

/*
Functions make it possible to encapsulate some logic and reuse it in our code
The basics of working with functions in Go will look familiar to you if you have
have worked with functions in other languages
*/
func main() { // main is the entry point of the Go program
	fmt.Println("Hello World")

	printMessage("i am printMessage function") // we added the value to the variable and output the result.
	// If we have added parameters, we cannot use the function without them.
	// printMessage()// not enough arguments in call to printMessage
	//have ()
	//want (string)
	print("exit") // exiting and fmt.Println(m) will not be executed because we use the return keyword

	message, sun := printAndReturnMessageAndSum("i am printAndReturnMessageAndSum function", 5)
	fmt.Println(message, sun) // we print the result of the function
	/*
		Add two variables because they need to be initialized with
		the appropriate values
		we can skip parameters with _
		message, _ := printAndReturnMessageAndSum("i am printAndReturnMessageAndSum function", 5)
		fmt.Println(message)
		output i am printAndReturnMessageAndSum function
		before they can be used in calculations or other operations
	*/

	PrintUser(User{name: "walter"})
	ad := addTop(1, 2, 3, 4, 5)
	fmt.Println(ad) // output 15 because we add 1+2+3+4+5

	PrintTypesFuncExample()

	MakeSortExample()

	MakeReturnFunctionExample()

	MakeExampleMemorize()

	MakeExampleClosures()

	MakeExampleClosuresSimple()

	MakeHardClosureExample()

	MakeMethodsExample()

	MakeExampleNoPointers()

	MakeExamplePointers()

	MakeExamplePointers2()
	MakeSimpleExamplePointers()

	MakeExamplePointers3()

	MakeExampleSwitch()

	MakeExampleSwitch1()

	MakeExampleSimpleLoops()

	// MakeExampleDefer()
}

/*
Since the main function does not accept
does not take parameters and does not return any values, let's look at an example of functions that do this
*/

func printMessage(m string) { // we add parameters to our function
	fmt.Println(m) //Input parameters are enclosed in parentheses and separated by commas
	// name|m Type|string,
}

/*
A function declaration includes four
elements:
the func keyword,
the function name,
the list of input parameters,
and the type of return value
*but we can skip the list of input parameters and the type of return value*
*/

func printAndReturnMessage(m string) string { // we add parameters to our function
	fmt.Println(m) // The type of the return value is written after the parentheses
	//parentheses with the list of input parameters,
	//before the opening curly brace of the function body

	return m // for return value we use the return keyword like in other languages
}

/*
In a function that does not return a value, the operator
return operator is required only if it is necessary to exit the function before the last line.
*/

func print(m string) {
	if m == "exit" {
		fmt.Println("exiting")
		return
	}
	fmt.Println(m)

}

func printAndReturnMessageAndSum(m string, n int) (string, int) { // we can add many parameters in return value

	sum := Sum(n, n) // we can use functions in another functions
	return m, sum    // we can't return values without declaring types in return value
}

func Sum(a int, b int) int {
	return a + b
}

type User struct {
	name string
	id   uint
	age  byte
}

// we also can use structs in parameters
func PrintUser(user User) {
	messageForName := fmt.Sprintf("name: %s", user.name)

	fmt.Println(messageForName)

}

/*

Go also allows you to use variational parameters.
The variable parameter must be the last parameter

(name ...Type) and this value is a slice of Type

*/
// Example

func addTop(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v // += is the same as sum = sum + v
	}
	return sum
}

/*
it is good practice in go
to give at most 2 variables and at most 2 return values per function.
*/

func addTop2(a int, b int) (int, int) {
	return a + b, a - b
}

// we can also add the name to the return value
func addTop3(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return sum, diff
	/*
		If your function returns values, never use an empty
		return statement. This can make it very difficult to understand
		exactly what value it returns
	*/
}

// but usually we declared only types in return value

// Function types

type myFunc func(int, int) int // we declare a function type

var f = map[string]myFunc{
	"sum": func(a int, b int) int {
		return a + b
	},
	"sub": func(a int, b int) int {
		return a - b
	},
}

var a, b int = 5, 1

var SumResult = f["sum"](a, b)
var SubResult = f["sub"](a, b)

// In the examples
// we define a map f containing functions that can be called by a key (string)
//We call a function stored in the map by passing it arguments

// Anonymous functions

func PrintTypesFuncExample() {
	fmt.Println("------")
	fmt.Println("PrintTypesFuncExample")
	fmt.Println(SumResult)
	fmt.Println(SubResult)

	func() {
		fmt.Println("anonymous function: 1")
	}() // we use () to call the anonymous function in this function now

	// another example

	for i := 0; i < 10; i++ {
		func(j int) {
			fmt.Println("printing", j, "from inside anonymous function")
		}(i)
	}
	/*
		We pass the variable i inside the for loop
		to the anonymous function.
		It is assigned to the input parameter j of the anonymous function
		function
	*/
}

/*
Usually we often use anonymous functions in goroutines,
which we will talk about later on
*/

// We can declare a function as a parameter in another function

func customSort(data []int, less func(i, j int) bool) {
	/*
		Takes a less function that compares
		two slice elements and returns true
		if the first element should be before the second.
	*/
	sort.Slice(data, less) // we import Slice from the sort package
	//sorts a slice of data using the less function
}

func MakeSortExample() {
	fmt.Println("------")
	fmt.Println("MakeSortExample")
	data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	customSort(data, func(i, j int) bool {
		return data[i] < data[j]
	})
	fmt.Println(data)
}

// return Function from another function

func makeW(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func MakeReturnFunctionExample() {
	fmt.Println("------")
	fmt.Println("MakeReturnFunctionExample")
	fmt.Println(makeW(5)(3))
}

// we use short circuit

func memorize(f func(int) int) func(int) int {
	mem := make(map[int]int)
	return func(x int) int {
		if v, ok := mem[x]; ok {
			return v
		}
		result := f(x)
		mem[x] = result
		return result
	}
}

// memoize returns
// a closure function that caches the results of calculations of function f

func MakeExampleMemorize() {
	fmt.Println("------")
	fmt.Println("MakeExampleMemorize")
	square := func(x int) int {
		return x * x
	}

	square = memorize(square)
	fmt.Println(square(5))
	fmt.Println(square(5)) // cached
	fmt.Println(square(10))
}

/*
Closures are a powerful tool in Go
that allow you to create functions with saved state
and control variables from the surrounding context.
They make code more flexible and convenient for various tasks,
such as creating counters, filters, and caching functions.
*/

// let's take a closer look at the closure functions

// example

func MakeExampleClosures() {
	fmt.Println("------")
	fmt.Println("MakeExampleClosures")
	inc := func() func() int {
		i := 0
		return func() int {
			i++
			return i
		}
	}
	nextInt := inc()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	nextInt2 := inc()
	fmt.Println(nextInt2())
	fmt.Println(nextInt2())
	fmt.Println(nextInt2())
}

// simple example
// returns a short-circuit function that increments and returns the increment value
func incrementor() func() int {
	inc := 0            // local variable to be captured by the closure
	return func() int { // return an anonymous function
		inc++      // increase inc at each function call
		return inc // return the current value of inc
	}
}

func MakeExampleClosuresSimple() {
	fmt.Println("------")
	fmt.Println("MakeExampleClosuresSimple")
	inc := incrementor() // inc is function
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}

// Hard example

func makefilter(f func(int) bool) func([]int) []int {
	return func(data []int) []int {
		result := []int{}
		for _, v := range data {
			if f(v) {
				result = append(result, v)
			}
		}
		return result
	}
}

func MakeHardClosureExample() {
	fmt.Println("------")
	fmt.Println("MakeHardClosureExample")
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ifEven := makefilter(func(i int) bool {
		return i%2 == 0
	})
	fmt.Println(ifEven(data))

	ifOdd := makefilter(func(i int) bool {
		return i%2 != 0
	})
	fmt.Println(ifOdd(data))
}

/********
### Methods ###
*********/

// we can supplement user types with methods

type Person struct {
	name     string
	age      byte
	verified bool
}

// we need add methods in same package as type (about packages we'll talk about later)
func (p Person) IsAdult() (string, bool) { // method receiver is specified
	return p.name + " is adult =", p.age >= 18
}

// we declared receiver  like this func (p Person)
// usually receiver is the type name abbreviation like this: (p Person)

// how we can call methods
func MakeMethodsExample() {
	fmt.Println("------")
	fmt.Println("MakeMethodsExample")
	p := Person{
		name: "walter",
		age:  25, // add 13 and watch the output
	}
	p.Isverified()
	if !p.verified {
		fmt.Println("unverified person")
		return
	}
	fmt.Println(p.IsAdult())
}

func (p Person) Isverified() bool {
	if p.age > 18 {
		return true
	}
	return false
}

//### Summary ###
/* the question arises:
when to use a function and when to use a method?
The key factor here is the dependence of the function on other data
In all those cases where the logic depends on values,
values that are configured at startup or changed during program execution.
program, store these values in the structure and implement the logic as a
method
*/

/********
### Pointers ###
*********/

type person struct {
	balance int
	status  string
	id      uint
	name    string
	age     byte
}

var user = person{
	name:    "Gustavo",
	age:     35,
	balance: 1000,
}

// because we are not using pointer golang make a copy of the struct
func (p person) addTaxCopy(amount int) {
	p.balance -= amount // we make a copy of the sctruct and we change the balance of copy struct
	fmt.Println("person need to pay ->", amount)
}

func MakeExampleNoPointers() {
	fmt.Println("------")
	fmt.Println("MakeExamplePointers")
	user.addTaxCopy(100)
	fmt.Println("new balance -> ", user.balance) // if we run this we will get same balance
	// because we are not using pointer go make a copy of the struct
}

// lets use pointers

func (p *person) addTax(amount int) { // we use * to use pointer to the struct
	// we take address of the user where he stored
	p.balance -= amount
	fmt.Println("person need to pay ->", amount)
}

func MakeExamplePointers() {
	fmt.Println("------")
	fmt.Println("MakeExamplePointers")
	user.addTax(100)
	fmt.Println("new balance -> ", user.balance)

	alexey.addTax(200)
	fmt.Println("new balance -> ", alexey.balance)
	// and now we changed the balance of the user and we can see the change

	// exampleuser.addtax(100) // this will not work
	//  because add tax only works with person struct

	// exampleuser.addTax(100) also does not work
	// type User2 struct {
	// 	balance  int
	// 	name     string
	// 	age      byte
	// 	verified bool
	// }

	// var exampleuser = User2{
	// 	name: "Nikola",
	// 	age:  25,
	// }

}

// if we declared the struct with index

var alexey = &person{
	name:    "Alexey",
	age:     35,
	balance: 1000,
} // we got same result as user but why?

//because go automatically converts user to &user to call the addTax method.

//alexey is already a pointer to person (*person), so the method is called directly

/*
When you have a person structure value, Go can automatically
get its pointer (&user) to call a method with a pointer receiver (*person).
Thus, the use of:

&person (pointer) and
just person (value)
gives the same results, because Go automatically handles conversions
between pointers and values to match the method's receiver type.
So you see the same behavior when calling
the addTax method for both user and alexey.

*/

// another example

type BankAccount struct {
	owner   string
	balance int
}

// method to deposit money into the account
func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Printf("Deposit of %d to account %s \n", amount, b.owner)
}

// method to withdraw money from the account
func (b *BankAccount) Widthdraw(amount int) bool {
	if b.balance >= amount {
		b.balance -= amount // balance = balance - amount
		fmt.Printf("Withdraw of %d from account %s \n", amount, b.owner)
		return true

	}
	fmt.Printf("%s attempted to withdraw %d. Insufficient balance: %d \n", b.owner, amount, b.balance)

	return false
}

func NewBankAccount(owner string, initialBalance int) *BankAccount {
	account := BankAccount{owner: owner, balance: initialBalance}
	return &account
}

func (b *BankAccount) CheckBalance() string {

	return fmt.Sprintf("The balance of %s account is: %d\n", b.owner, b.balance)

}

func MakeExamplePointers2() {
	Alexey := NewBankAccount("Alexey", 1000)
	Alexey.Deposit(500)
	Alexey.Widthdraw(1000)

	fmt.Println(Alexey)
}

/*The & operator is used to get the address of a variable.
When you use & before a variable, you get a pointer to that variable.
*/

func MakeSimpleExamplePointers() {
	x := 10                                      // 0xc4242
	p := &x                                      // p is 0xc4242
	fmt.Println("Address of x (value of p):", p) // address of x like

	fmt.Println("Value of x through pointer p:", *p) // value of x like 10
	*p = 20                                          // Change value by address
	fmt.Println("Adress of p:", &p)

	fmt.Println(x) // 20
}

/*
Memory Address | Value
---------------|------
0xc4242        | 10   (x)
0xc4243        | 0xc4242 (p) (address of x)

after *p = 20

Memory Address | Value
---------------|------
0xc4242        | 20   (x)
0xc4243        |  0xc4242 (p) (address of x)

The variable x and the pointer p occupy different memory locations



*/

type database struct {
	id   uint
	user string
}

type Server struct {
	db *database //This allows Server to manipulate
	// database data through a pointer, which is efficient for modifying data.
}

func (s *Server) GetUser(id uint) (string, error) { // we can return error to check
	if s.db == nil {
		return "", fmt.Errorf("user not found")
	}
	return s.db.user, nil
}

func MakeExamplePointers3() {
	fmt.Println("------")
	fmt.Println("MakeExamplePointers3")
	server := Server{
		db: &database{
			id:   1,
			user: "Alexey",
		},
	}
	user, _ := server.GetUser(1)

	fmt.Println(user)
}

/*
Why Use Pointers?
Efficiency: Passing a pointer to a method instead of a copy of the
struct is more efficient, especially if the struct is large.
Mutability: Methods with pointer receivers can modify the original
struct, allowing for changes to the state of the object.
*/

/********
### Switch ###
*********/

type OperationType int // bank operation type

const (
	Deposit      OperationType = iota // 0
	Withdraw                          // 1
	CheckBalance                      // 2

	/*
		iota is a powerful and concise way to create sequences of
		related constants in Go. Its automatic incrementing behavior
		makes your code more readable
		and less error-prone when defining enumerated constants or bit flags.
	*/
)

// Operate performs a bank operation based on the type provided
func (acc *BankAccount) Operate(optype OperationType, amount int) {
	switch optype { //  Switch statement to handle different operation types
	case Deposit: // If case is Deposit, call the Deposit method of BankAccount
		acc.Deposit(amount)
	case Withdraw: // If case is Withdraw, call the Withdraw method of BankAccount
		acc.Widthdraw(amount)
	case CheckBalance: // If case is CheckBalance, print the balance
		fmt.Println("Balance:", acc.balance)

	default: // Handle any invalid operation types
		fmt.Println("Invalid operation")
	}
}

func MakeExampleSwitch() {
	fmt.Println("------")
	fmt.Println("MakeExampleSwitch")
	acc := BankAccount{"Alexey", 1000}
	acc.Operate(Deposit, 500)
	acc.Operate(Withdraw, 1000)
	acc.Deposit(200)
	acc.Operate(CheckBalance, 0)
}

// another example

type ActionType int

const (
	MoveNorth   ActionType = iota // 0
	MoveSouth                     // 1
	MoveEast                      // 2
	MoveWest                      // 3
	CheckStatus                   // 4
) // define constants for different action types

type Player struct {
	X, Y   int
	Name   string
	Health uint
}

func NewPlayer(name string) *Player {
	return &Player{
		Name:   name,
		Health: 100,
	}
}

func (p *Player) Act(actionType ActionType) {
	switch actionType {
	case MoveNorth:
		p.Move(0, 1)
	case MoveSouth:
		p.Move(0, -1)
	case MoveEast:
		p.Move(1, 0)
	case MoveWest:
		p.Move(-1, 0)
	case CheckStatus:
		fmt.Println("Player name:", p.Name)
		fmt.Println("Player position:", p.X, p.Y)
		fmt.Println("Player health:", p.Health)
	default:
		fmt.Println("Invalid action")
	}
}
func (p *Player) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
	fmt.Printf("Player %s moved to (%d, %d)\n", p.Name, p.X, p.Y)
}

func MakeExampleSwitch1() {
	fmt.Println("------")
	fmt.Println("MakeExampleSwitch1")
	player := NewPlayer("Alexey")
	player.Act(MoveNorth)
	player.Act(MoveSouth)
	player.Act(MoveEast)
	player.Act(MoveWest)
	player.Act(CheckStatus)
}

/********
### if-else/for-loops/range/break ###
*********/

// It will not take much time, unlike other topics,
// as it is very simple here so

func exampleFor() {

	for i := 0; i < 10; i++ { // simple basic for structure in c like languages
		fmt.Println(i)
	}

	/*

		for init|i:=0; condition|i<10; post|i++ {
			// loop body
		}
	*/

	/*
		# Explanation:
		- `i := 0`: Initializes the loop variable `i` to 0.
		- `i < 10`: The loop continues as long as this condition is true.
		- `i++`: Increments `i` by 1 after each iteration.
		- `fmt.Println(i)`: Prints the current value of `i`.
	*/

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < len(numbers); i++ { // len is a function that returns the length of a slice
		fmt.Println(numbers[i])
	}

	// output 1 2 3 4 5 6 7 8 9
}

func exampleWhile() {

	i := 0 // initialize i outside the loop

	for i < 10 { // loop as long as i is less than 10 (similar to a while loop)
		fmt.Println(i)
		i++ // increment i after each iteration
	}

	/*
		# Explanation:
		- This `for` loop behaves like a `while` loop in other languages.
		- The loop continues as long as the condition `i < 10` is true.
		- `i++` increments `i` by 1 in each iteration.
	*/

}

func exampleInfiniteLoop() {

	i := 0

	for { // infinite loop (no condition)
		if i > 10 { // exit condition inside the loop
			break // break out of the loop if i is 10 or more
		}
		if i == 5 { // skip the rest of the loop if i is 5
			i++
			continue // continue to the next iteration
		}
		fmt.Println(i)
		i++
	}

	/*
		# Explanation:
		- This `for` loop runs indefinitely because there is no condition.
		- The `if` statement inside the loop checks if `i` is 10 or more.
		- `break` exits the loop when the condition is met.
		- `i++` increments `i` by 1 in each iteration.
	*/
}

func exampleRange() {
	nums := []int{2, 4, 6, 8, 10} // # Define a slice of integers
	for i, num := range nums {    // # Use `range` to iterate over the slice
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}

	nums1 := []string{"a", "b", "c"}
	for _, num := range nums1 { // we can skip indes if we want or skip the value
		if num == "b" {
			fmt.Println("Found b: break")
			break
		}
		fmt.Printf("Value: %s\n", num)
	}
	fmt.Println("break out of the loop")

	/*
		# Explanation:
		- `nums := []int{2, 4, 6, 8, 10}`: Creates a slice of integers.
		- `for i, num := range nums`: Iterates over the slice with `range`.
		- `i` is the index and `num` is the value at that index.
		- `fmt.Printf("Index: %d, Value: %d\n", i, num)`: Prints the index and value.
	*/

	// # Example with a map
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	} // map is don't sort by default so
	// every time it will be random
	for name, age := range ages { // # Use `range` to iterate over the map
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}

	/*
		# Explanation:
		- `ages := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 35}`: Creates a map.
		- `for name, age := range ages`: Iterates over the map with `range`.
		- `name` is the key and `age` is the value.
		- `fmt.Printf("Name: %s, Age: %d\n", name, age)`: Prints the key and value.
	*/
}

func exampleIfElse(score int) {
	if score >= 90 { // if score is 90 or above its an a
		fmt.Println("A")
	} else if score >= 80 { // if score is 80 or above its a b
		fmt.Println("B")
	} else if score >= 70 { // if score is 70 or above its a c
		fmt.Println("C")
	} else if score >= 60 { // if score is 60 or above its a d
		fmt.Println("D")
	} else { // if score is below 60 its an f
		fmt.Println("F")
	}

	/*
		# Explanation:
		- The `if` statement checks the condition `score >= 90`.
		- If the condition is true, it executes the code block inside the `if`.
		- If the condition is false, it checks the next `else if` condition, and so on.
		- The `else` block runs if none of the previous conditions are true.
	*/

}

// another example

func MakeExampleSimpleLoops() {
	fmt.Println("------")
	fmt.Println("MakeExampleSimpleLoops")
	name := []string{"Alice", "Niko", "Charlie"}
	for _, nam := range name {
		if nam == "Niko" {
			fmt.Println("Found Niko")
			break
		}
		if nam == "Alice" {
			continue
		}
		fmt.Println(nam)
	}
}

/********
### Defer/Panic ###
*********/

func printMessages() {
	var messages = []string{
		"Hello",
		"World",
	}
	messages[2] = "Bye"   // panic: runtime error: index out of range [2] with length 2
	fmt.Println(messages) //
	// we can call panic function to stop the program
	panic("something went wrong")
}

func MakeExampleDefer() {
	fmt.Println("------")
	fmt.Println("MakeExampleDefer")
	printMessages()
	/*
		First, it can be used to postpone the execution of multiple closures
		within a function. These closures are executed according to the
		"last in - first out" principle, i.e. the defer operator specified
		last is executed first.
		last in - first out" principle, i.e.
		the defer operator specified last is executed first

	*/
	defer func() { // defer calls in end of all functions

		// defer = function postponement
		// Recover from a panic and print the error message
		// some overhead 50ms
		handlePunic()
	}()

	// Simulate a critical error situation
	err := fmt.Errorf("Critical error occurred")

	// Trigger a panic
	panic(err)

	// This line will not be executed
	fmt.Println("Program will not reach here")

	/* # Explanation:
	- defer # The defer statement allows us to make function call(s) that
	will be executed just before the function exits, regardless of how it exits.
	- recover # The recover function is used to recover from a panic
	and resume normal execution. It returns the value passed to the panic call.
	- fmt.Errorf # This constructs a new error message.
	- panic # This triggers a panic containing the specified error, halts the program's normal execution, and the deferred function is then executed to recover from the panic.

	Remember that using panic should be reserved for truly exceptional situations where the program cannot safely continue. It is considered a best practice to handle errors gracefully without resorting to panics whenever possible.
	*/
}

func handlePunic() {
	if r := recover(); r != nil { // recover is standart function of go
		fmt.Println("Recovered from panic:", r)
	}
}

// some real example using this

func DoSomething(ctx context.Context, db *sql.DB, v1 string, v2 string) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil { // check error in golang is a good practice
		return err
	}
	defer func() {
		if err == nil {
			err = tx.Commit()
		}
		if err != nil {
			tx.Rollback()
		}
	}()
	_, err = tx.ExecContext(ctx, "INSERT INTO table VALUES (?, ?)", v1, v2)
	if err != nil {
		return fmt.Errorf("insert failed: %w", err)
	}
	// do something
	return nil
}

// in go we usually use error as return value to check errors

func MakeExampleError() error {
	fmt.Println("------")
	fmt.Println("MakeExampleError")
	err := DoSomething(context.Background(), nil, "v1", "v2")
	if err != nil {
		fmt.Println("Error:", err)
		return fmt.Errorf("DoSomething failed: %w", err)
	}

	return nil
}
