package main
import "fmt"

func add(x int, y int) (z1 int, z2 int){

	defer fmt.Println("Hello")
	z1 = x + y
	z2 = x - y
	fmt.Println("Before Return")
	return 
}

// Function as Argument

func test2(Myfunc func(x int) int){
	fmt.Println(Myfunc(7))
}

//Return type as aFunction

func test3(x string) func(){
	return func ()  {
		fmt.Println(x)
	}
}


func main(){

	ans1, ans2 := add(15,10)
	fmt.Println(ans1, ans2)

	//Inner function
	test := func (x int) int {
		return x * -1
	}

	fmt.Println(test(7))
	test2(test)

	test3("Hello")()


}