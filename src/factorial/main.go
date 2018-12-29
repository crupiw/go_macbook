package main

import (
	"fmt"
)

func Factorial(i int) int {
	if i <= 1{
		return 1
	}
	return i * Factorial(i-1)
}
func main(){
	fmt.Println("factorial 5 is ::",Factorial(5))
	TowersOfHanoi(3)
}
func TOHUtil(num int, from string, to string, temp string){
	if num < 1 {
		return
	}
	TOHUtil(num-1, from, temp,to)
	fmt.Println("Move disk ",num,"from peg",from, " to peg",to)
	TOHUtil(num-1,temp,to,from)
}
func TowersOfHanoi(num int){
	fmt.Println("the Sequence of moves involved in the Tower of Hanoi are:")
	TOHUtil(num,"A","C","B")
}