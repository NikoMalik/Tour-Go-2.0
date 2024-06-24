package main

import (
	"fmt"
	"io"
)

// abstract type in go
// which describes the behavior of other types
type IntArray interface { // we use type to declare user type
	// interface is nil by default

	// describe what interface can do

	GET() ([]int, error) // set of methods

	POST() ([]int, error)

	PUT() ([]int, error)
}

// if interface is nil it will be panic

// in go we usually named interfaces with er

// like io.Reader
// like io.Writer
// fmt.Stringer
// io.Closer
// etc

type Logic interface {
	Process(data string) string
}

// func (f *Client) Process(data string) string {
// 	return f.Process(data)
// }

type Client struct {
	L Logic
	io.Closer
}

func (c *Client) Program() error {
	result := c.L.Process("Niko is here")
	fmt.Println(result)
	return nil
}

// Interfaces enable polymorphism in Go, allowing different types to be treated uniformly based on the behaviors they provide.

// SimpleLogic is a concrete type that implements the Logic interface
type SimpleLogic struct{}

// Process method for SimpleLogic type that satisfies the Logic interface
func (sl *SimpleLogic) Process(data string) string {
	return "Processed: " + data
}

// DummyCloser is a dummy implementation of io.Closer for demonstration purposes
type DummyCloser struct{}

// Close method for DummyCloser to satisfy the io.Closer interface
func (dc *DummyCloser) Close() error {
	fmt.Println("Closed!")
	return nil
}

func main() {
	logic := &SimpleLogic{}
	client := &Client{
		L:      logic,
		Closer: &DummyCloser{},
	}
	client.Program()

	fmt.Println("------")

	Mike := &Employee{
		name: "Mike",
		sum:  1000,
	}

	Ceo := &CEO{
		name: "Ceo",
		sum:  20000,
	}

	function(Mike, 200, "add")

	function(Ceo, 100, "sub")

	fmt.Println(Ceo.sum)

	fmt.Println(Mike.sum)

	fmt.Println("------")

	apiServer := &ApiServer{
		storer: &MongoDB{},
	}

	result, err := apiServer.storer.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	apiServerSql := &ApiServer{
		storer: &MySQL{},
	}

	result, err = apiServerSql.storer.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	fmt.Println("------")

	quackers := []Quacker{
		&Duck{}, &Duckling{}, &Drake{},
	}

	for i := 0; i < len(quackers); i++ {
		quackers[i].Quack()
	}

	s := &QuackerType{
		quack: &Drake{},
	}

	s.quack.Quack()

	// or use range if you want this is simple

}

// another example more simple

type Money interface {
	Add(amount float64) float64
	Sub(amount float64) float64
}

type Employee struct {
	name string
	sum  float64
}

type CEO struct {
	name string
	sum  float64
}

func (e *Employee) Add(amount float64) float64 { // our method add
	e.sum += amount * 1.2
	return e.sum
}

func (e *Employee) Sub(amount float64) float64 { // our method sub
	e.sum -= amount / 2
	return e.sum
}

func (c *CEO) Add(amount float64) float64 { // our method add
	c.sum += amount * 2
	return c.sum
}

// we use interfaces to simply implement polymorphism in go
/*
Interfaces are used as types for variables, function parameters,
and return values to provide flexible and reusable code components.
*/

func (c *CEO) Sub(amount float64) float64 { // our method sub
	c.sum -= amount
	return c.sum
}

func function(money Money, amount float64, operation string) float64 {
	switch operation {
	case "add":
		return money.Add(amount)
	case "sub":
		return money.Sub(amount)
	default:
		return 0
	}

}

// another example

type Storer interface {
	Get() ([]int, error)

	Put(int) ([]int, error)
}

type MongoDB struct {
}

type MySQL struct {
}

func (m *MongoDB) Get() ([]int, error) {
	return []int{1, 2, 3}, nil // nil == null remember nil is nothing
}

func (m *MongoDB) Put(number int) ([]int, error) {
	fmt.Println("Put", number)
	return []int{1, 2, 3}, nil
}

func (m *MySQL) Get() ([]int, error) {
	return []int{4, 5, 6}, nil
}

func (m *MySQL) Put(number int) ([]int, error) {
	fmt.Println("Put", number)
	return []int{4, 5, 6}, nil
}

type ApiServer struct {
	storer Storer
}

// another example simple to understand

type Quacker interface {
	Quack()
}

type Duck struct {
}

type Drake struct{}

type Duckling struct{}

func (d *Duck) Quack() {

	fmt.Println("duck")
}

func (d *Drake) Quack() {
	fmt.Println("drake")
}

func (d *Duckling) Quack() {
	fmt.Println("quackling")
}

// also we can make this

type QuackerType struct {
	quack Quacker
}
