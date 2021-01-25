package main

import (
	"example.com/hello/protocol/first"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	//p := &first.Person{
	//	Name:  "张三",
	//	Id:    0,
	//	Email: "123@qq.com",
	//	Phones: []*first.Person_PhoneNumber{
	//		{
	//			Number: "18753938432",
	//			Type:   0,
	//		},
	//	},
	//	LastUpdated: nil,
	//}

	//fmt.Println(p.String())

	//book := &first.AddressBook{}
	// ...

	// Write the new address book back to disk.
	//out, err := proto.Marshal(p)
	//if err != nil {
	//	log.Fatalln("Failed to encode address book:", err)
	//}
	//if err := ioutil.WriteFile("addressbook", out, 0644); err != nil {
	//	log.Fatalln("Failed to write address book:", err)
	//}

	// Read the existing address book.
	in, err := ioutil.ReadFile("addressbook")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book := &first.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	fmt.Println(book.String())
}
