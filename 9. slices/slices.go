package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("unint:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	// s[3] = "d" => panic: runtime error: index out of range [3] with length 3
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	s[3] = "z"
	fmt.Println("apdupd:", s)
	fmt.Println("len:", len(s), "cap:", cap(s)) // cap is doubled (cap = cap x 2) when it reaches the limit

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c, "len c:", len(c), "cap c:", cap(c))

	l := s[2:5] // 2 to 4
	fmt.Println("sl1:", l)

	l = s[:5] // 0 to 4
	fmt.Println("sl2:", l)

	l = s[2:] // 2 to end
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"j", "k", "l"}
	if slices.Equal(t, t2) {
		fmt.Println("t and t2 are equal")
	} else {
		fmt.Println("t and t2 are not equal")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
