package main
import "fmt"
func IncrementPassByPointer(ptr *int) {
	(*ptr)++
}
func max(x,y int) int{
	if x > y {
		return x
	}
	return y
}
type student struct {
	rollNo int
	name string
}
func main() {
	s:= "hello, world!"
	r:= []rune(s)
	r[0]= 'H'
	s2 := string(r)
	fmt.Println(s2)
        numbers := []int{1,2,3,4,5,6,7,8,9,10}
        sum := 0
	for i:=0 ; i < len(numbers);i++ {
		sum += numbers[i]
	}
	fmt.Println("Sum is ::", sum)
	sum = 0
 	for index, val:= range numbers{
		sum += val
		fmt.Print("[", index,",",val,"]")
	}
	fmt.Println("Sum is ::", sum)
	kvs := map[int]string{1:"apple",2:"banana"}
	for k,v := range kvs {
		fmt.Println(k,"->",v)
	} 
	str := "Hello, World!"
	for index, c := range str{
		fmt.Print("[",index,",",string(c),"]")
	}
	fmt.Println()
	fmt.Println(max(10,20));
	data := 10
	ptr  := &data
	fmt.Println("Value stored at variable var is ", data)
	fmt.Println("Value stored at variable var is ", *ptr)
	fmt.Println("the address of variable var is ", &data)
	fmt.Println("the address of variable var is ", ptr)
	i:= 10
	fmt.Println("Value of i before increment is: ", i)
	IncrementPassByPointer(&i)
	fmt.Println("Value of i after increment is: ", i)
        stud:= student{1,"Johnny"}
	fmt.Println(stud)
	fmt.Println("Student name::",stud.name)
	pstud := &stud
	fmt.Println("Student name::",pstud.name)
	fmt.Println(student{rollNo:2,name:"Ann"})
	fmt.Println(student{name:"Ann",rollNo:2})
	fmt.Println(student{name:"Alice"})
}
