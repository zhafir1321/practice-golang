package main

import (
	"fmt"
	"math"
	"reflect"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000
	fmt.Println(reflect.TypeOf(n))

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(reflect.TypeOf(d))

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
	fmt.Println(reflect.TypeOf(math.Sin(n)))
}

// const => A const statement can appear anywhere a var statement can. (Outside of a function, at the package level.)
// const => Constant expressions perform arithmetic with arbitrary precision.
