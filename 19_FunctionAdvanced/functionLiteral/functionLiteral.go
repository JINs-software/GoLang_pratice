package main

import "fmt"

// 1) 별칭으로 함수 정의 줄여쓰기 (opFunc)
type opFunc func(a, b int) int

func getOperator(op string) opFunc {
	if op == "+" {
		return func(a, b int) int {
			// 함수 리터럴을 사용하여 더하기 함수를 정의하고 반환
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	fn := getOperator("*") // fn은 함수 타입 변수
	result := fn(3, 4)     // 함수 타입 변수를 사용하여 함수를 호출
	fmt.Println(result)

	// 아래와 같이 직접 호출도 가능
	result = func(a, b int) int {
		return a + b
	}(3, 4)
}
