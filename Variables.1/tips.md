# So which of these styles is best to use?

As with anything else,
use whichever way makes your intentions as clear as possible.
intentions. In most cases, inside functions, it is better to declare variables using the `:=` operator. Outside of functions, use declaration lists
in those rare cases when you need to declare several variables at once at the package level.
at the package level.
In some cases it will be better to refrain from using the `:=` operator inside functions.
functions.

- When you need to initialize a variable with a null value, use the form `var x int`. By doing so, you will clearly show that you wanted to create a variable with a null value.
  variable with a null value

- When an untyped constant or literal is assigned to a variable and that constant or literal defaults to a different type than the variable should have.
  type that the variable should have, use the long form of the var declaration with a type statement. Nothing prevents you from using the := operator, though, by specifying the type of the variable using a type conversion:
  `x := byte(20)`, the idiomatic approach boils down to writing this
  as `var x byte = 20`

constants in Go are computed at compile time and do not produce

> Surprisingly, the Go compiler <mark style="background: #BBFABBA6;"> allows you to create unused constants</mark> using the const keyword. This is because
> any side effects. This makes it easy to remove them: if a constant is not used anywhere, it is simply not included in the compiled
> binary
