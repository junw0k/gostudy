package main

import (
	"fmt"
	"time"
)

func Task(id int) {
	fmt.Printf("작업 [%d] 시작->", id)
	fmt.Printf("작업 [%d] 완료\n", id)
}

func main() {

	for i := 0; i < 10; i++ {
		go Task(i + 1)

	}
	time.Sleep(time.Second)
	fmt.Println("메인 함수 종료")
}
