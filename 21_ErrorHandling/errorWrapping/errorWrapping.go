package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func MulipleFromString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str)) // (1) 스캐너 생성, 한 단어씩 읽는 bufio 패키지의 Scanner 생성
	scanner.Split(bufio.ScanWords)                      // (2) 한 단어씩 끊어 읽기, Scanner 객체의 Split() 메서드를 통해 어떻게 끊을지 명시

	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		// (6) 에러 감싸기, readNextInt() 발생 에러를 감싸서 pos 정보를 에러에 추가
		return 0, fmt.Errorf("failed to readNextInt(), pos:%d err:%w", pos, err)
	}

	pos += n + 1
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("failed to readNextInt(), pos:%d err:%w", pos, err)
	}

	return a * b, nil
}

// 다음 단어를 읽어 숫자로 변환하여 반환
// 변환된 숫자, 읽은 글자 수, 에러를 반환
func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() { // (3) Scan 메서드를 통해 첫번째 단어 읽기
		return 0, 0, fmt.Errorf("Failed to scan")
	}
	word := scanner.Text()
	number, err := strconv.Atoi(word) // (4) 문자열을 int 타입으로 변경(숫자가 아닌 문자가 섞여 있는 경우 NumberError 타입의 에러 반환)
	if err != nil {
		// (5) 에러 감싸기
		return 0, 0, fmt.Errorf("failed to convert word to int, word:%s err:%w", word, err)
	}
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MulipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		// (7) 감싸진 에러를 다시 꺼내올 땐 errors 패키지의 As() 함수를 이용
		if errors.As(err, &numError) { // 감싸진 에러가 NumError인지 확인
			fmt.Println("NumberError:", numError)
		}
	}
}

func main() {
	readEq("123 3")
	readEq("123 abc")
}
