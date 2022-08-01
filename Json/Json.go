package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Name   string `json:"name"`
	Author Author `json:"author"`
}

type Author struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}



func main() {

	author := Author{Name: "Rabindranath Tagore", Age : 56, Address: "Bangalore"}
	book := Book{Name: "I too had a lv story", Author: author}
	byteArray, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))
	var book1 Book

	err1 := json.Unmarshal(byteArray,&book1)

	if(err1 != nil){
		fmt.Println(err1)
	}

	fmt.Printf("%+v\n",book1)
}
