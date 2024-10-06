package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Age int
}

// *Student 타입은 String 메서드를 포함함
func (s *Student) String() string {
	// 인터페이스 변수를 구체화된 타입으로 타입 변환을 하려면
	// 해당 타입이 인터페이스 메서드 집합을 포함하고 있어야 함
	return fmt.Sprintf("Student Age: %d", s.Age)
}

type Actor struct {
}

func (a *Actor) String() string {
	return "Actor"
}

func PrintAge(stringer Stringer) {
	// *Student 타입으로 변환
	// String() 메서드를 포함한 것은 *Student, Student로의 타입 변환은 불가
	s := stringer.(*Student)

	fmt.Printf("Age: %d\n", s.Age)
}

func StudentConvert(stringer Stringer) {
	student := stringer.(*Student)
	fmt.Println(student)

	// 이 함수에 *Actor 타입 인스턴스 전달 시 런타임 에러 발생
	// stringer 인터페이스 변수가 *Actor 타입 인스턴스를 가리키고 있기 떄문이다.
	// => stringer 변수는 내부에 *Actor 인스턴스를 가리키고 있다.
}

func main() {
	s := &Student{15} // *Student 타입 변수 s 선언 및 초기화
	PrintAge(s)       // 변수 s를 인터페이스 인수로 PrintAge() 함수 호출

	actor := &Actor{}
	StudentConvert(actor) // <- 런타임 에러 발생
}
