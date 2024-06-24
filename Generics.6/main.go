package main

import "fmt"

/********
### Generics ###
*********/
// we use it rarely in go but you need to know it

// type generalization
type Custom[K comparable, V any] struct { // we use generics
	data map[K]V
} // https://go.dev/blog/comparable

/*

comparable indicates that the generalized type K must be comparable.
This means that values of type K must be comparable to each other
using comparison operations (e.g. ==, !=, <, >, etc.).


any is any type in go



*/

func (c *Custom[K, V]) Get(key K, value V) error {
	c.data[key] = value
	return nil
}

func NewCustomMap[K comparable, V any]() *Custom[K, V] {
	return &Custom[K, V]{
		data: make(map[K]V),
	}
}

// we can use also generics in functions
var a = []int{1, 2, 3}
var b = []byte{1, 2, 3}

func Sum[Y int | byte](input []Y) Y {
	var result Y

	for i := 0; i < len(input); i++ {
		result += input[i]
	}
	return result
}

// another example

func searchElement[K comparable](elements []K, element K) bool {
	for i := 0; i < len(elements); i++ {
		if elements[i] == element {
			return true
		}
	}
	return false
}

func main() {
	m := NewCustomMap[string, int]()
	m.Get("map: m", 1)
	fmt.Println(m)
	m1 := NewCustomMap[int, string]()
	m1.Get(1, "map: m1")
	fmt.Println(m1)

	m2 := NewCustomMap[int, int]()
	m2.Get(1, 1)
	fmt.Println(m2)
	fmt.Println("------")

	fmt.Println(Sum(a))

	fmt.Println("------")

	fmt.Println(searchElement(a, 1))

}
