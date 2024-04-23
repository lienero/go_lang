package main

import (
	"fmt"
	"sort"
)


type Student struct {
	Name string
	Age int
}

// []Student의 별칭 타입 Students
type Students []Student

// 구조체 슬라이스 정렬
// Sort 함수를 이용하기 위해서는 Len(), Less(), Swap() 세 메서드가 필요함
func (s Students) Len() int {return len(s)}
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age } // 나이 비교
func (s Students) Swap(i, j int) { s[i], s[j] = s[j], s[i]}

func main() {
	s := []Student{{"화랑", 31}, {"백두산", 52}, {"류", 42}, {"켄", 38}, {"송하나", 18}}

	sort.Sort((Students(s))) // 정렬
	fmt.Println(s)
}