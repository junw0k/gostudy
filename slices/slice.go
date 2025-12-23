package main

//slice

import (
	"fmt"
)

func main() {
	var s []string
	fmt.Println("uninit:", s, s == nil, "len:", len(s), "cap:", cap(s))

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	fmt.Println("s[0]:", s[0])

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	fmt.Println("s", s)

	s = append(s, "e", "f")
	fmt.Println("s", s)

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println("cpy", c)

	l := s[:5]
	fmt.Println("sl2:", l)

	twoD := make([][]int, 3)
	for i := range 3 {
		ilen := i + 1
		twoD[i] = make([]int, ilen)
		for j := range ilen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
