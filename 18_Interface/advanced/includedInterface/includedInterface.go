package includedinterface

// 구조체에서 다른 구조체를 구조체를 포함된 필드로 가질 수 있듯
// 인터페이스도 다른 인터페이스를 포함할 수 있음
type Reader interface {
	Read() (n int, err error)
	Close() error
}

type Writer interface {
	Write() (n int, err error)
	Close() error
}

type ReadWriter interface {
	Reader // Reader 메서드 집합 포함
	Writer // Writer 메서드 집합 포함
	// Close() error는 같은 메서드 형식이므로 합쳐져서 하나의 Close() 메서드로 포함됨
}
