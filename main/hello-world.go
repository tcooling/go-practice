package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("Hello, world!")
	values()
	variables()
	constants()
	forLoop()
	ifElse()
	switchConditional()
}

func values() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func variables() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// variables declared without value are zero-valued
	// int is 0, string is empty, bool is false etc.
	var e int
	fmt.Println(e)

	// this is same as declaring a var, only available inside functions
	// var f string = "apple"
	f := "apple"
	fmt.Println(f)
}

var aa int = 1

const s string = "constant"

func constants() {
	// Go supports constants of char, str, bool and numeric values
	fmt.Println(s, aa)

	// A numeric constant has no type until it is given one
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n)) // Sin expects a float64
}

func forLoop() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1 // can do i++ or i += 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			// use continue to go to next iteration of loop
			continue
		}
		fmt.Println(n)
	}
}

func ifElse() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// A statement, e.g. setting 9, can precede an if
	// it will be available in all branches
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

func switchConditional() {
	i := 2
	fmt.Println("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it is the weekend")
	default:
		fmt.Println("it is a weekday")
	}

	// switch without an expression can replicate if/else logic
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it is before noon")
	default:
		fmt.Println("it is after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	fmt.Println("Hello in language:", getGreeting("es"))

	// Cannot assign the result of a switch statement directly to a variable
	// return exits the func, not the switch, below does not work
	// var a = switch 1 {
	// case 1:
	// 	return "1"
	// case 2:
	// 	return "2"
	// default:
	// 	return "default"
	// }
}

// Example of switch in func
func getGreeting(language string) string {
	switch language {
	case "en":
		return "Hello"
	case "es":
		return "Hola"
	case "fr":
		return "Bonjour"
	default:
		return "Hi"
	}
}
