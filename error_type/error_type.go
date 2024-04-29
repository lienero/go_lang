package main

import "fmt"

type PasswordError struct { // 에러 구조체 선언
	Len int
	RequireLen int
}

func (err PasswordError) Error() string { // Erro() 메서드 선언, error 인터페이스로 사용가능
	return "암호 길이가 짧습니다."
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{ len(password), 8} // error 반환
	}

	return nil
}

func main() {
	err := RegisterAccount("myID", "myPw") // ID, PW 입력
	if err != nil { // 에러 확인
		// errInfo, ok := err.(PasswordError); : 초기문
		// 초기문 : 사용할 변수를 초기화할 때 주로 사용
		// ok : 조건문
		if errInfo, ok := err.(PasswordError); ok { // 인터페이스 변환
			fmt.Printf("%v Len:%d RequireLen:%d\n", errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("회원 가입됐습니다.")
	}
}