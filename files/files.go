package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func checkError(err error){
	if(err != nil){
		panic(err)
	}
}

func main() {

	content := "Hi this is Pavan!!"
	file, err := os.Create("./MyFile.txt")

	checkError(err)
	length, err := io.WriteString(file,content)
	checkError(err)

	defer file.Close()

	fmt.Println("Length of the File is:",length)
	readFile("./MyFile.txt")


}

func readFile(FileName string){
	dataByte, err := ioutil.ReadFile(FileName)
	checkError(err)
	fmt.Println(string(dataByte))
	
}