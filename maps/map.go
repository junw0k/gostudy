package main

// dictionary 형태
// assciative_data type
import (
	"fmt"
)

func main() {
	m := make(map[string]int) // make 내장함수

	m["elem1"] = 7
	m["elem2"] = 4

	fmt.Println("map:", m)
	fmt.Println("len : ", len(m))

	v1 := m["elem1"]
	fmt.Println("v1", v1)

	v3 := m["elem3"]

	fmt.Println("v3:", v3) // 존재안할시에 0 반환
	fmt.Println("map:", m)

	fmt.Println("len : ", len(m))

	delete(m, "elem2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

}
