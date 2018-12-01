package main

import (
	"fmt"
	"geometry/mygeo"
)

func main() {
        np := mygeo.Point{1,2}
	pptr := &np	
        pptr.ScaleBy(2)
	fmt.Println(np)
	perim := mygeo.Path{
		{1,1},
		{5,1},
		{5,4},
		{1,1},
	}
	p := mygeo.Point{4, 9}
	q := mygeo.Point{2, 6}
        r := &mygeo.Point{1,2}
	r.ScaleBy(10)
        
	fmt.Println(&r)
	fmt.Println(*r)
        fmt.Println(r)
	fmt.Println()

	fmt.Println(mygeo.Distance(p, q))
	fmt.Println(p.Distance(q))
        fmt.Println(perim.Distance())
}
