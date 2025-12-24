package main

import (
	"fmt"
)

func makezero(ival int) {
	ival = 0
}

func ptrmakezero(iptr *int) {
	*iptr = 0
}

func set1array(arr []int) {
	for i := range len(arr) {
		arr[i] = 1
	}
}

func set0array(arr []int) {
	for i := range len(arr) {
		arr[i] = 0
	}
}

func setptr(arr *[]int) {
	//모든 요소 1로 반영
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = 1
	}
	fmt.Println("함수 내부 원본 수정:", *arr)
}

// 포인터가 없는 경우:
func attemptAppend(s []int, val int) {
	s = append(s, val) // 복사된 이정표만 4칸으로 늘어남. 메인문의 원본 이정표는 여전히 3칸.
}

// 포인터가 있는 경우
func realAppend(s *[]int, val int) {
	*s = append(*s, val) // 원본 이정표의 길이를 4로 바꾸고 주소도 새로 갈아 끼움 -> 원본 반영 O
}

func main() {
	i := 1
	fmt.Println(i)

	makezero(i)
	fmt.Println("pass by value ", i)

	ptrmakezero(&i)
	fmt.Println("pass by address ", i)

	fmt.Println("---array, slices, address----")

	var s []int
	s = make([]int, 3)

	fmt.Println("s:", s)
	fmt.Println("&s", &s)
	fmt.Println("&s[0]", &s[0])
	fmt.Println("&s[1]:", &s[1])

	fmt.Println("--set1array--")
	set1array(s)
	fmt.Println(s)

	fmt.Println("--set0array--")
	set0array(s)
	fmt.Println(s)

	setptr(&s)

	fmt.Println("[]으로 넘기고, append")
	attemptAppend(s, 100)
	fmt.Println(s)

	fmt.Println("address, append ")
	realAppend(&s, 100)
	fmt.Println(s)

	// slice 구조 struct {cap, len, 요소포안터}
	// 해당 구조를 바꾸는경우 len 값등.. ->ptr
}
