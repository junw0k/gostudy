package main

import (
	"errors"
	"fmt"
)

type Account struct {
	balance int
}

// 1. 원본 수정을 위해 포인터(*) 수신자 사용
// 2. 에러 전달을 위해 반환 타입에 error 추가
func (a1 *Account) Withdraw(amount int) error {
	if amount > a1.balance {
		// 생성된 에러를 호출한 곳으로 던져줍니다.
		return errors.New("잔액이 부족합니다")
	}
	// 차감
	a1.balance -= amount
	return nil // 에러가 없으면 nil(비어있음)을 반환
}

func main() {
	myAcc := Account{balance: 1000}

	err := myAcc.Withdraw(1500)

	if err != nil {
		// 에러가 있다면 출력하고 종료
		fmt.Println("실패:", err)
	} else {
		// 에러가 nil(없음)이라면 성공 출력
		fmt.Println("출금 성공! 현재 잔액:", myAcc.balance)
	}
}
