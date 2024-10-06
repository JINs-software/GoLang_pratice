package main

// interface{}는 메서드를 가지고 있지 않은 빈 인터페이스
// 모든 타입이 빈 인터페이스로 쓰일 수 있음
// 빈 인터페이스는 어떤 값이든 받을 수 있는 함수, 메서드, 변수값을 만들 때 사용함

import "fmt"

func PrintVal(v interface{}) { // 빈 인터페이스를 인수로 받는 함수
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float64 %f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		fmt.Printf("Not supported type: %T: %v\n", t, t)
	}
}

type Student struct {
	Age int
}

func main() {
	PrintVal(10)      // int
	PrintVal(3.14)    // float64
	PrintVal("Hello") // string
	PrintVal(Student{15})
}
