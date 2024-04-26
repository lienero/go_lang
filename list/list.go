package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 리스트 : 불연속된 메모리에 데이터를 저장
	v := list.New() // 새로운 리스트 생성
	e4 := v.PushBack(4) // 리스트 뒤에 요소 추가
	e1 := v.PushFront(1) // 리스트 앞에 요소 추가
	v.InsertBefore(3, e4) // e4 요소 앞에 요소 삽입
	v.InsertAfter(2, e1) // e1 요소 뒤에 요소 삽입

	// Front() 메서드는 가장 첫 번째 요소를 반
	// Next() 현재 요소의 다음 요소를 반환합니다.
	for e := v.Front(); e != nil; e = e.Next() { // 각 요소 순회
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
	// Back() : 메서드는 가장 마지막 요소를 반환
	// Prev() : 현재 요소의 이전 요소를 반환
	for e := v.Back(); e != nil; e = e.Prev() { // 각 요소 역순 순회
		fmt.Print(e.Value, " ")
	}
}