package main

import (
	"fmt"
	"log"

	proto "github.com/golang/protobuf/proto"
)

func main() {

	author := &Author{
		Id:   5,
		Name: "Dostoyevski",
		Books: []*Book{
			{Id: 1, Name: "Kumarbaz", Type: BookType_PAPER},
			{Id: 2, Name: "Beyaz Geceler", Type: BookType_PAPER},
			{Id: 3, Name: "Yeraltindan Notlar", Type: BookType_EBOOK},
		},
	}

	data, err := proto.Marshal(author)
	if err != nil {
		log.Fatal("Marshal error: ", err)
	}

	// print the raw data
	fmt.Printf("Raw: %v \nRaw: %q \n", data, data)
	fmt.Println("-------")

	newAuthor := &Author{}

	err = proto.Unmarshal(data, newAuthor)
	if err != nil {
		log.Fatal("Unmarshal error: ", err)
	}

	fmt.Printf("Yazar ID: %d \nYazar ADI: %s\n", newAuthor.GetId(), newAuthor.GetName())
	fmt.Println("Kitaplari: ")

	for _, book := range newAuthor.GetBooks() {

		fmt.Printf("Kitap ID: %d, Kitap ADI: %s, ", book.GetId(), book.GetName())
		fmt.Printf("Kitap TIP: %s\n", book.GetType())
	}

}
