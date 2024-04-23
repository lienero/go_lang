package main

import "fmt"

type account struct {
	balance int
}

// 일반 함수 표현
func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

// 메서드 표현 : 리시버를 사용 (리시버 : 메서드가 속하는 타입을 알려주는 기법)
func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func main() {
	// balance가 100인 account 포인터 변수 생성
	a := &account{ 100 }
	withdrawFunc(a, 30) // 함수 형태 호출
	a.withdrawMethod(30) // 메서드 형태 호출
	fmt.Printf("%d \n", a.balance)
}