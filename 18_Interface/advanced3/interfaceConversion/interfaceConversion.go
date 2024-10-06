package main

import "fmt"

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {
}

func (f *File) Read() {
	fmt.Println("File's Read()..")
}
func (f *File) Close() {
	fmt.Println("File's Close()..")
}

func ReadFile(reader Reader) {
	// Reader 인터페이스 변수를 Closer 인터페이스로 타입 변환
	closer := reader.(Closer)
	closer.Close()
}

func main() {
	file := &File{}
	ReadFile(file)
}
