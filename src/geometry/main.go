package main

import (
	"fmt"
	"geometry/mygeo"
)

func main() {
	p := mygeo.Point{1, 2}
	q := mygeo.Point{4, 6}
	fmt.Println(mygeo.Distance(p, q))
	fmt.Println(p.Distance(q))
}
