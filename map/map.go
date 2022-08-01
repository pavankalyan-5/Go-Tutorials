package main
import "fmt"

func main(){
	var mp map[string]int = map[string]int{
		"Apple" : 1,
		"Orange" : 2,
	}
	fmt.Println(mp["Apple"])
	mp["Apple"] = 100
	// delete(mp,"Apple")  Delete from map
	fmt.Println(mp)

	vegetables := map[string]int{
		"Tomato": 20,
		"Onion" : 50,
	}
	fmt.Println(vegetables)
	val , ok := vegetables["Tomato"]
	if ok == true{
		fmt.Println(val)
	}

	for key,val := range vegetables{
		fmt.Println(key,val)
	}

	

}
