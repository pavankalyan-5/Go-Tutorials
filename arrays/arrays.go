package main
import "fmt"

func main(){

	arr := [5]int{1,2,3,4,5}
	sum:=0

	for i:=0;i<len(arr);i++{
		sum+=arr[i]
	}
	fmt.Println(sum)

	//Slicing of the array

	var x[] int = arr[1:3]
	fmt.Println(x)
	fmt.Println("Length of the array: ",len(x))
	fmt.Println("Capacity of the array: ",cap(x))
	b := append(x, 10)
	fmt.Println(b)

	var a []int = [] int{11,12,13,14,15}
	fmt.Println(a)

	for i,ele := range a{                // i -> Index
		fmt.Println(i,ele)				//  ele -> element	
	}

}