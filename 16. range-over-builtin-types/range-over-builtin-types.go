package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		fmt.Println("Index:", i, "Number:", num)
		sum += num
	}
	fmt.Println("Sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("Index:", i)
			fmt.Println(num == 3)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana", "c": "cherry"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("Key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
