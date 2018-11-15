package main

import "fmt"

func add(args ...int) int {
	total := 0
	for _, x := range args {
		total += x
	}
	return total
}

func main() {
	fmt.Println(add(1, 2, 3))
	y := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(add(y...))
}
