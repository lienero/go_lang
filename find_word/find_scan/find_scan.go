package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LineInfo struct {
	lineNo int
	line   string
}

type FindInfo struct {
	filename string
	lines    []LineInfo
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다. ex) ./find_scan word filepath")
		return
	}

	// 실행 인수 가져오기
	word := os.Args[1] // 찾으려는 단어
	files := os.Args[2:]
	findInfos := []FindInfo{}
	for _, path := range files {
		// 파일 찾기
		findInfos = append(findInfos, FindWordInAllFiles(word, path)...)
	}

	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		fmt.Println("-----------------------")
		for _, lineinfo := range findInfo.lines {
			fmt.Println("\t", lineinfo.lineNo, "\t", lineinfo.line)
			fmt.Println("-----------------------")
			fmt.Println()
		}
	}
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path) // 파일 경로에 해당하는 파일 목록을 가져옴
}

func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}
	filelist, err := GetFileList(path) // 파일 리스트 가져오기
	if err != nil {
		fmt.Println("파일 경로가 잘못되었습니다. err:", err, "path:", path)
		return findInfos
	}
	for _, filename := range filelist { // 각 파일 별로 검색
		findInfos = append(findInfos, FindWordInFiles(word, filename))
	}
	return findInfos
}

func FindWordInFiles(word, filename string) FindInfo {
	findInfo := FindInfo{filename, []LineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다.", filename)
		return findInfo
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file) // 스캐너를 만듭니다.
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) { // 한 줄씩 읽으면 단어 포함 여부 검색
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	return findInfo
}