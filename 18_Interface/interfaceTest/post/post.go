package post

import "fmt"

type PostSender struct {
}

func (k *PostSender) Send(parcel string) {
	fmt.Printf("Post sends %v parcel\n", parcel)
}
