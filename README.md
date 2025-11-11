# go-practice
Learning Go

Following the [Go by example](https://gobyexample.com/) tutorials.

## Running Instructions

Run the program:
```shell
go run hello-world.go
```

Package it into a binary:
```shell
go build hello-world.go
```

Execute the binary:
```shell
./hello-world
```

## Notes

- Filenames use dashes `hello-world.go`
- You can omit the return type if the function doesn't return anything
- `...someVar` means 0 or more values
- Boolean logic is similar to Scala, `&&`/`||`/`!`
- `[...]` in array definition means compiler will count elements
- Trailing commas at end of array defs etc
- Check arrays empty by checking against `nil`
- Can do similar to Scala extension methods

## Questions

- do not understand `make([]string, 3, 5)` third param?
    - is capacity underlying array size?
- if a map is 0 valued, how to save a 0?

## To Read

- https://go.dev/tour/list
- https://go.dev/blog/slices-intro