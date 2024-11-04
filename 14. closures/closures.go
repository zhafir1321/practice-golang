package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func counter() func() (int, func()) {
	count := 0

	return func() (int, func()) {
		count++
		reset := func() {
			count = 0
		}
		return count, reset
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	nextCount := counter()

	val1, _ := nextCount()
	fmt.Println("Counter:", val1)

	val2, _ := nextCount()
	fmt.Println("Counter:", val2)

	val3, resetFunc3 := nextCount()
	fmt.Println("Counter:", val3)

	resetFunc3()

	val4, _ := nextCount()
	fmt.Println("Counter after reset:", val4)
}
