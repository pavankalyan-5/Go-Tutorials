package main 
import "fmt"

func main(){

	age := 17

	if age >= 18 {
		fmt.Println("You can Ride Alone!!")
	} else if age >= 14 {
		fmt.Println("Ride with Parents")
	}else{
		fmt.Println("You can't ride")
	}
}