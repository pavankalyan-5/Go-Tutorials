package main

import (
	"errors"
	"fmt"
)

func sumOf(start, end int) (int, error) {

	if start > end {
		return 0 , errors.New("Start should not be greater than end")
	}
	total := 0
	for i := start; i <= end; i++ {
		total += i
	}
	return total,nil
}

func main() {

	ans ,err := sumOf(12,10)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(ans)
	}
}