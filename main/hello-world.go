package main

import (
	"fmt"
	"maps"
	"math"
	"slices"
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
	arrays()
	slicesFunc()
	mapFunc()
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

func arrays() {
	// In Go, an array is a numbered sequence of elements of specific length
	// Slices are more common
	var a [5]int
	fmt.Println("emp:", a) // array is 0 valued (0, empty string etc)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// below does not compile as 5 is out of bounds
	// fmt.Println("get:", a[5])

	fmt.Println("len:", len(a))

	// declare and init on one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// have the compiler count the elements
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// if specify ':', elements between are zeroed
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	c := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("c:", c)

	// compose arrays for multi dimensional
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d:", twoD)
}

func slicesFunc() {
	// slices are more powerful than arrays
	// an unitialised slice has length of 0

	var s []string
	fmt.Println("un-init:", s, s == nil, len(s) == 0)

	// to create a slice with a non zero length, use make
	// this will fill it with 0 valued (empty string) values
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	// slices support append, need to assign to variable
	// as could get back new slice value
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// get slice of elements between 2 and 5 (excluding 5)
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}

	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}

func mapFunc() {
	// similar to dicts or hash maps etc
	// map[key-type]val-type
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	// second return value, prs, indicates if key is present
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
