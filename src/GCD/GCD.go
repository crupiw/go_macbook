package main

import "fmt"

func GCD(m int , n int) int {
	if m < n {
		return GCD(n,m)
	}
	if m%n == 0 {
		return n
	}
	return GCD(n,m%n)
}
func fibonacci (n int ) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
func Permutation(data[]int, i int, length int){
	if length == i {
		PrintSlice(data)
		return
	}
	for j := i; j < length ; j++ {
		swap(data,i,j)
		Permutation(data,i+1,length)
		swap(data,i,j)
	}
}
func BinarySearchRecursive(data[] int,low int, high int, value int) int {
	mid := low + (high-low)/2
	if data[mid] == value {
		return mid
	} else if data[mid] < value {
		return BinarySearchRecursive(data, mid+1, high, value)
	} else {
		return BinarySearchRecursive(data, low, mid-1, value)
	}
}

func swap(data[] int, x int,y int) {
	data[x], data[y] = data[y],data[x]
}
func PrintSlice(data [] int) {
	fmt.Printf("%v:: len= %d cap= %d \n", data, len(data),cap(data))
}
func main() {
	var data [5] int
	for i := 0; i < 5; i++ {
		data[i] = i
	}
	bindata := [6]int{9, 22, 31, 55, 65,105}
	fmt.Println(GCD(1000,20))
	fmt.Println(fibonacci(10))
	Permutation(data[:],0,5)
	fmt.Println(BinarySearchRecursive(bindata[:],0, len(data),55))
}
