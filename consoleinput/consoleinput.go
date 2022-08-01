package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"

)

func main(){
	scanner := bufio.NewScanner(os.Stdin) // for scanning input
	fmt.Printf("Enter your birth year: ")
	scanner.Scan()
	age, _ := strconv.ParseInt(scanner.Text(),10,64) // (value , base , bit)
	fmt.Printf("Your age will be %d by the end of 2022", 2022-age)
}