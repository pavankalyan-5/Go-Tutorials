package main

import "fmt"

func main(){
	x := 0
	for x < 5{
		fmt.Println(x)
		x+=1
	}

	for x:=0;x<=5;x+=1{
		fmt.Println(x)
	}
}