package main

import "fmt"

func main() {
	// 값을 저장할 변수 선언
	var a int
	var b int

	// 입력 두개 받기
	// n은 성공적으로 입력한 개수이고, err은 입려기 발생한 에러를 반환
	// 표준 입력
	// n, err := fmt.Scan(&a, &b)
	// 서식에 맞춘 입력
	// n, err := fmt.Scanf("%d %d\n", &a, &b)
	// 엔터 키로 입력을 종료해야함
	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}