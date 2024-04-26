package main

import "fmt"

func sum(nums ...int) int { // 가변 인수를 받는 함수
	sum := 0

	fmt.Printf("nums 타입: %T\n", nums) // nums 타입 출력
	for _,v := range nums {
		sum += v
	}
	return sum
}

// func Print(args ...interface{}) string { // 모든 타입을 받는 가변 인수
// 	for _, arg := range args { // 모든 인수 순회
// 		switch f := arg.(type) { // 인수의 타입에 따른 동작
// 		case bool:
// 			val := arg.(bool) // 인터페이스 변환
// 		case float64:
// 			val := arg.(float64)
// 		case int:
// 			val := arg.(int)
// 		}
// 	}
// 	return val
// }

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5)) // 인수 5개를 사용합니다.
	fmt.Println(sum(10, 20)) // 인수 2개를 사용합니다.
	fmt.Println(sum()) // 인수 0개를 사용합니다.
}