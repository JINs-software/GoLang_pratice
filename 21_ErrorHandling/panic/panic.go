package main

import "fmt"

func divide(a, b int) {
	if b == 0 {
		panic("b는 0일 수 없습니다")
	}

	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func main() {
	divide(9, 3)
	divide(9, 0) // panic 발생
	//main.divide(0x9?, 0x3?)
	//    C:/Users/HurJin/0. GameServerStudy/3. Go/GoLang_pratice/21_ErrorHandling/panic/panic.go:7 +0xee
	//main.main()
	//    C:/Users/HurJin/0. GameServerStudy/3. Go/GoLang_pratice/21_ErrorHandling/panic/panic.go:15 +0x29
}
