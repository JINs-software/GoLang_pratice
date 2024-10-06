package main

import "fmt"

func f() {
	fmt.Println("f() start")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic 복구 - ", r)
		}
	}()

	g() // <- g()로 부터 패닉 전파, 아래 fmt.Println은 실행되지 않고, defer의 함수 리터럴을 통해 복구됨
	fmt.Println("f() end")
}

func g() {
	fmt.Printf("9 / 3 = %d\n", h(9, 3))
	fmt.Printf("9 / 0 = %d\n", h(9, 0))
}

func h(a, b int) int {
	if b == 0 {
		panic("divide by 0") // 패닉 발생
	}
	return a / b
}

func main() {
	f()
	fmt.Println("main end")
}
