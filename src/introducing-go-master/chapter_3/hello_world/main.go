package main

import "fmt"

func main() {
	var x string = "Hello World"
	const y string = "Hello World 2"
	var (
		a = 5
		b = 2
		c = 3
	)
	z := 23
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(z)
	name := "Max"
	fmt.Println("My dog's name is", name)
}
