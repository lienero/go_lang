package main

import "fmt"

func main() {
	// := (var 밑 변수 선언문을 생략)
	a := [5]int{1, 2, 3, 4, 5} // 배열 선언
	b := [5]int{500, 400, 300, 200, 100}

	// range는 배열의 각 요소를 순회하면서 인덱스와 요솟값 두 값을 반환
	for i, v := range a { // i : 인덱스, v: 원솟값
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	b = a // a 배열을 b 변수에 복사
	fmt.Println()
	for i, v := range b { // i : 인덱스, v: 원솟값
		fmt.Printf("b[%d] = %d\n", i, v)
	}
}