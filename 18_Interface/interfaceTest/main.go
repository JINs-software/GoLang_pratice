package main

import (
	"18_Interface/interfaceTest/fedex"
	"18_Interface/interfaceTest/post"
)

// Sender 인터페이스를 만듬
type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	postSender := &post.PostSender{}
	SendBook("어린 왕자", postSender)

	fedexSender := &fedex.FedexSender{}
	SendBook("인간 실격", fedexSender)
}
