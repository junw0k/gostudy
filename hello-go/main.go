package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name  string
	Score int
}

func checkPass(score int) string {
	if score > 80 {
		return "합격"
	}

	return "불합격"
}

func calculateArea(redius float64) float64 {
	return math.Pow(redius, 2) * math.Pi
}

func main() {
	fmt.Println("hello GO!")
	r := 5.0
	result := calculateArea(r)
	fmt.Printf("반지름이 %.1f일 때 원의 넓이: %.2f\n", r, result)

	students := []Student{
		{Name: "철수", Score: 85},
		{Name: "영희", Score: 65},
		{Name: "민수", Score: 95},
	}

	fmt.Println("----성적---")
	for _, s := range students {
		result := checkPass(s.Score)
		fmt.Printf("이름: %s | 점수 : %d | 결과 :%s\n", s.Name, s.Score, result)
	}
}
