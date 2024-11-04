package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = "initial"
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))

	var b, c int = 1, 2
	fmt.Println(b, c)
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))

	var d = true
	fmt.Println(d)
	fmt.Println(reflect.TypeOf(d))

	var e int
	fmt.Println(e)
	fmt.Println(reflect.TypeOf(e))

	f := "apple"
	fmt.Println(f)
	fmt.Println(reflect.TypeOf(f))
}

/*
	:= syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.
*/
