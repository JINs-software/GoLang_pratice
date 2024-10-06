package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err // 에러 발생 시 에러 반환
	}
	defer file.Close() // 함수 종료 전 파일 닫기

	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n') // 파일에서 한 줄 읽기
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, line) // 파일에 문자열 쓰기
	return err
}

const filename string = "data.txt"

func main() {
	line, err := ReadFile(filename) // 파일 읽기 시도
	if err != nil {
		err = WriteFile(filename, "This is WriteFile") // 파일 생성
		if err != nil {
			fmt.Println("파일 생성 실패 ", err)
			return
		}

		// 다시 읽기 시도
		line, err = ReadFile(filename)
		if err != nil {
			fmt.Println("파일 읽기 실패 ", err)
			return
		}
	}

	fmt.Println("파일 내용: ", line)
}
