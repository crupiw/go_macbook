package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}
func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, x2)
	w := distance(x1, y1, x1, x2)
	return l * w
}
func circleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}
func circleAreaPtr(c *Circle) float64 {
	return math.Pi * c.r * c.r
}
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x2, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y2)
	return l * w
}
func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5
	c := Circle{0, 0, 5}
	cPtr := Circle{0, 0, 9}
	cc := new(Circle)
	cc.r = 6
	cc.x = 0
	cc.y = 0
	fmt.Println("Rectangle Area")
	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println("Circle Area")
	fmt.Println(circleArea(cx, cy, cr))
	fmt.Println("Circle values")
	fmt.Println(c.x, c.y, c.r)
	fmt.Println("Circle Area struct")
	fmt.Println(circleArea(cc.x, cc.y, cc.r))
	fmt.Println("Circle area pointer")
	fmt.Println(circleAreaPtr(&cPtr))
	fmt.Println("Circle area method")
	fmt.Println(c.area())
	r := Rectangle{0, 0, 11, 12}
	fmt.Println("Rectangle area method")
	fmt.Print(r.area())
}
