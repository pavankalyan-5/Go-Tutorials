package main
import "fmt"

type Point struct{
	x int32
	y int32
}

type Circle struct{
	radius float32
	Origin *Point
}

type Student struct{

	Name string
	Grades []int
	age int 
}
func (s Student) getAge() int {
	return s.age
}

func (s *Student) setAge(x int) {
	s.age = x
}

func (s Student) getAverageGrade() float32{
	g := s.Grades
	sum := 0
	for i := 0;i<len(g);i+=1  {
		sum += g[i]
	}
	return float32(sum) / float32(len(g))
}
func main(){
	var p1 Point = Point{1,2}
	var p2 Point = Point{2,3}
	fmt.Println(p1 ,p2)

	p3 := &Point{3,5}
	c1 := Circle{4.75,p3}
	fmt.Println(c1.Origin)

	s := Student{"Pavan",[]int {1,2,3,4},21}
	fmt.Println(s)
	fmt.Println(s.getAge())
	s.setAge(17)
	fmt.Println(s.getAge())

	fmt.Println(s.getAverageGrade())
}
