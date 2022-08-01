package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
const URL = "https://lco.dev/"
func main() {
	response, err := http.Get(URL)
	if err != nil{
		panic(err)
	}
	fmt.Printf("Type of the response is: %T\n",response)

	defer response.Body.Close()

	dataByte , err := ioutil.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}
	content := string(dataByte)
	fmt.Println(content)

}