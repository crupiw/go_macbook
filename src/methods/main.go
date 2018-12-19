package main
import ("fmt"
	"math"
       )
type Shape interface {
	Area() float64
	Perimeter() float64
}
type Circle struct{
	radius float64
}

type MyInt int

func (data MyInt)increment1(){
	data = data + 1
}
func (data *MyInt) increment2(){
	*data = *data + 1
}
type Rect struct {
	width  float64
	height float64
}

func(r Rect) Area() float64 {
	return r.width * r.height
}
func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}
func (c Circle) Area()float64{
	return math.Pi*c.radius*c.radius
}
func (c Circle)Perimeter()float64{
	return 2 * math.Pi*c.radius
}
func TotalPerimeter(shapes ...Shape)float64{
	var peri float64
	for _,s:= range shapes{
		peri+=s.Perimeter()
	}
	return peri
}
func TotalArea(shapes ...Shape)float64{
	var area float64
	for _,s:= range shapes{
		area+=s.Area()
	}
	return area
}
func main() {
	r:= Rect{width: 10, height: 10}
	c:= Circle{radius:10}
	fmt.Println("Area: ", r.Area())
	fmt.Println("Perimeter:",r.Perimeter())
	ptr:=&Rect{width:10, height:5}
	fmt.Println("Area: ", ptr.Area())
	fmt.Println("Perimeter:",ptr.Perimeter())
	var data MyInt = 1
	fmt.Println("value before increment1()call:",data)
	data.increment1()
	fmt.Println("value after increment1()call:",data)
	data.increment2()
	fmt.Println("value after increment2()call:",data)
	fmt.Println("Total Area:", TotalArea(r,c))
	fmt.Println("Total Perimeter:",TotalPerimeter(r,c))
}
	
