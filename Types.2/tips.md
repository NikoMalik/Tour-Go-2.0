# So when is it best to use each of of these types?

### In most cases and most often the usual _int_ is used (99%)

but in other cases like `this` :

- If you are working with binary files or a network protocol,
  If you are working with binary files or a network protocol that uses integers of a certain size or sign, select the appropriate integer type
  appropriate integer type

- If you are writing a library function that must work with any integer type, write two functions, one of which will use int64 type for parameters and variables, and the second - uint64 type.
  integer type, write two functions, one of which will use the int64 type for parameters and variables, and the second one will use the uint64 type.
  (We'll talk more about functions and their parameters later)

- Scenario in which you have a table in your database with a lot of entries with an integer for an id, which is always positive. If you store this data as an int one bit of every entry is effectively useless and when you scale this, you are losing a lot of space when you could have just used a uint and saved it. Similar scenario can be thought of while transmitting data, transmitting tons of integers to be precise. Also, uint has double the range for positive integers compared to their counterpart signed integers due to the extra bit, so it will take you longer to run out of numbers. Storage is cheap now so people generally ignore this supposedly minor gain.

The other usecase is type-safety. A uint can never be negative so if a part of your code is delicate to negative numbers, it can prove to be pretty handy. It's better to get the error before wasting resource on the data just to find out it's impermissible because it's negative.

- RGB images uint8 for 8-bit images, uint16 for 10/12/16 bit.

>

# Useful tips on the topic

- If you use [...], you get an array. The slice is formed using [ ].

- Functions such as len are built into the Go language because the actions they perform cannot be accomplished with functions written in the Go language.
  actions they perform cannot be performed by functions written by a programmer.
  by a programmer. As you've seen, the len function can take as input
  any array or slice. A little later we will see that it can also work with strings and maps. Trying to pass a variable of any other type to the len function will result in a compile-time error. Go does not allow developers to write functions that can to behave in the same way

- Use structures to organize data with different types of fields. For example:
  ```go
      type Person struct {
      Name string
      Age  int
  }
  ```

---

```go
var x = []int{}

```

- This code creates a zero-length slice that is not nil (comparing it to the value nil returns false). In all other
  in all other respects, a zero-length slice behaves exactly the same as a slice equal to
  The only case when such a zero-length slice may be needed is when converting a slice to a nil-length slice.
  is to convert the slice to JSON format. We will talk about
  more about this later

- If the slice is used as a buffer , it is better to specify a non-zero length

- In the make function call, it's better to specify zero
  length and non-zero capacity. This allows you to add elements using the append function. If the actual number of elements is less than
  the specified capacity, you will not get extra zero values at the end of the
  of the slice. If the number of elements exceeds the specified capacity, your code won't
  will not panic.

- If you specified a slice length in the make function call, before you
  use the append function, make sure it is exactly what you want. Otherwise, you may end up with unnecessary null values at the beginning of the slice

- Try not to modify slices after
  create a slice based on them or after they have been obtained by slicing

- <https://www.youtube.com/watch?v=Tl7mi9QmLns> about maps

- The comma-ok idiom is used in Go when you want to distinguish between reading a value and returning a null value
  value
