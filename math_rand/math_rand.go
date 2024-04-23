package main

import (
	"bufio" // io를 담당하는 패키지
	"fmt"
	"math/rand" // 랜덤값을 생성하는 패키지
	"os"        // 표준 입출력 등을 가지고 있는 패키지
	"time"      // 날짜, 시각, 시간 등을 다루는 패키지
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n) // int 타입값을 입력받음
	if err != nil {
		stdin.ReadString('\n') // 에러 바ㅣㄹ생 시 입력 스트림을 비움
	}
	return n, err
}

func main() {
	rand.Seed(time.Now().UnixNano()) // 시간 값을 랜덤 시드로 설정

	r := rand.Intn(100) // 0~99 사이 랜덤값 생성
	cnt := 1

	for {
		fmt.Print("숫자값을 입력하세요:")
		n, err := InputIntValue() // 숫자값 입력
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else {
			if n > r { // t숫자값 비교
				fmt.Println("입력하신 숫자가 더 큽니다")
			} else if n < r {
				fmt.Println(("입력하신 숫자가 더 작습니다."))
			} else {
				fmt.Println("숫자를 맞췄습니다. 축하합니다. 시도한 횟수:", cnt)
				break // 같을 경우 메세지를 출력하고 break로 종료
			}
			cnt++
		}
	}
}
