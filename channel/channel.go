package main

import "fmt"

func Sum(a, b int, resultChan chan int) {
	resultChan <- a + b
}

func main() {
	done := make(chan bool)

	go func() {
		fmt.Println("일하는중")
		done <- true
	}()
	<-done
	fmt.Println("일끝")

	resultchan := make(chan int)
	go Sum(10, 20, resultchan)
	answer := <-resultchan
	fmt.Println(answer)

}
