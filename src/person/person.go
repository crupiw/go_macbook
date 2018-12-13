package main

import "fmt"

type Person struct {
	Name string
}
type Android struct {
	Person
	Model string
}

type Shape interface {
	area() float64
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func main () {
	a := new(Person)
	a.Name = "Bill Crupi"
	a.Talk()
	b := new(Android)
	b.Name = "Bob Crupi"
	b.Person.Talk()
}
