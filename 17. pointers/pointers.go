package main

import "fmt"

func zeroVal(ival int) { // ival is a copy of the original i
	ival = 0 // this change will not affect the original i
}

func zeroPtr(iptr *int) { // iptr is a pointer to the original i
	*iptr = 0 // this change will affect the original is
}

func main() {

	i := 1
	fmt.Println("Initial:", i)

	zeroVal(i)                 // passing the value of i
	fmt.Println("zeroVal:", i) // i is still 1

	zeroPtr(&i)                // passing the address of i
	fmt.Println("zeroPtr:", i) // i is now 0

	fmt.Println("Pointer:", &i) // prints the address of i
}
