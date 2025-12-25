//interface

package main

import "fmt"

type Payer interface {
	pay(amount int)
}

type CreditCard struct {
	CardName string
}

func (C CreditCard) pay(amount int) {
	fmt.Printf("%s카드로 %d원 결제 완료\n", C.CardName, amount)
}

type Cash struct{}

func (cash Cash) pay(amount int) {
	fmt.Printf("현금으로 %d원 지불 완료\n", amount)
}

func ProcessPayment(p Payer, amount int) {
	p.pay(amount)
}

func main() {
	card := CreditCard{CardName: "삼성"}
	cash := Cash{}
	ProcessPayment(card, 5000)
	ProcessPayment(cash, 3000)
}
