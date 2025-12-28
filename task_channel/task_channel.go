package main

import (
	"fmt"
	"time"
)

func Task(id int, ch chan int) {
	time.Sleep(time.Millisecond * 10)
	ch <- id
}

func main() {
	// 출력을 메인에서 실행 고루틴 내부에서 print 직접 x
	resultChan := make(chan int)
	for i := 1; i <= 10; i++ {
		go Task(i, resultChan)
	}
	results := make([]int, 11)

	for i := 1; i <= 10; i++ {
		id := <-resultChan
		results[id] = id
	}
	for i := 1; i <= 10; i++ {
		fmt.Printf("작업 [%d] 완료 확인\n", results[i])
	}
}
