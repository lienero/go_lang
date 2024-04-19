package main

import (
	"fmt"
	"strings"
)

func ToUpper(str string) string {
	var builder strings.Builder // 내부에 슬라이스를 갖고 있기에 WriteRune 메소드를 통해 문자를 더할 때 메모리를 새로 생성하지 않고, 기존 메모리 공간에 더함
	for _, c := range str { // range 를 이용하여 한 글자씩 순회
		if c >= 'a' && c <= 'z' { // 소문자인지 확인
			builder.WriteRune('A' + (c - 'a')) // 대문자로 변경
		} else {
			builder.WriteRune(c) // 소문자가 아니면 그대로 추가
		}
	}

	return builder.String()
}

func main() {
	str := "hello World!"
	fmt.Println(ToUpper(str))
}