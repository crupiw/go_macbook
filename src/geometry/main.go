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
	fmt.Println()
	perim := mygeo.Path{
		{1,1},
		{5,1},
		{5,4},
		{1,1},
	}
	p := mygeo.Point{4, 9}
	fmt.Println()
	q := mygeo.Point{2, 6}
	fmt.Println()
        r := &mygeo.Point{1,2}
	r.ScaleBy(10)
        s := mygeo.Point{9,11}
        (&s).ScaleBy(3)
        fmt.Println(s)
	fmt.Println()
	fmt.Println(&r)
	fmt.Println(*r)
        fmt.Println(r)
	fmt.Println()

	fmt.Println(mygeo.Distance(p, q))
	fmt.Println(p.Distance(q))
        fmt.Println(perim.Distance())
        fmt.Println() 
	fmt.Println(pptr.Distance(q))
	fmt.Println()
        fmt.Println((*pptr).Distance(q))
}
