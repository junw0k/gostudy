package main

import "fmt"

type Account struct {
	Owner   string
	Balance int
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
}

func (a Account) GetBalance() int {
	return a.Balance
}

func main() {

	a1 := Account{"홍길동", 1000}
	fmt.Println(a1)
	a1.Deposit(1000)
	fmt.Println(a1)

	fmt.Println("balance : ", a1.GetBalance())

}
