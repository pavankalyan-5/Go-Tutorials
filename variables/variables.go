package main 
import "fmt"

func main(){
	var name string = "Pavan Kalyan"
	var age int = 21
	number := 142 // implicit 
	fmt.Println("Hello", name ,"of age" ,age)
	fmt.Printf("%T\n",number) // Type of the variable
	fmt.Printf("%v\n",number) // Value of the variable 
	var x string = fmt.Sprintf("Hello, %s of age %v",name,age)
	fmt.Println(x)
}