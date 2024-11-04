package main

import "fmt"

func plus(a int, b int) int { // When you have multiple consecutive parameters of the same type, you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
	return a + b

}

func plusPlus(a, b, c int) int { // When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
	return a + b + c
}

func main() {
	res := plus(1, 2) // Call a function just as you'd expect, with name(args).
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
