package main

import "fmt"

type PasswordError struct { // 에러 구조체 선언
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string { // Error() 메서드 정의
	return "암호 길이가 짧습니다."
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8}
	}
	return nil
}

func main() {
	err := RegisterAccount("myID", "123")
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok { // error 인터페이스에서 구조체 타입으로 변환
			fmt.Printf("%v Len: %d RequireLen: %d\n",
				errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("회원 가입 성공!")
	}
}
