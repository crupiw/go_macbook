package main
import "fmt"
func SumArray(data[]int)int {
	size:= len(data)
	total:= 0
	for index := 0; index <size; index++ {
		total = total+ data[index]
	}
	return total
}
func SequentialSearch(data[]int, value int)bool{
	size := len(data)
	for i := 0; i <size; i++ {
		if value == data[i] {
			return true
		}
	}
	return false
}
func BinarySearch(data[]int, value int) bool{
	size := len(data)
	var mid int
	low:=0
	high := size -1
	for low <= high {
		mid = low + (high-low)/2
		if data[mid] == value {
			return true
		} else {
			if data[mid] < value {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
		return false
}
// Rotating a list by K positions.
func RotateArray(data[]int,k int) {
	n := len(data)
	ReverseArray(data,0,k-1)
	ReverseArray(data,k,n-1)
	ReverseArray(data,0,n-1)
}
func ReverseArray(data []int, start int,end int){
	i:= start
	j:= end
	for i < j {
		data[i], data[j] = data[j],data[i]
		i++
		j--
	}
}
func MaxSubArraySum( data [] int) int {
   size := len( data)
   maxSoFar := 0
   maxEndingHere := 0
   for i := 0; i < size; i++ {
     maxEndingHere = maxEndingHere + data[i]
     if maxEndingHere < 0 {
     	maxEndingHere = 0
     	fmt.Println("MaxEndingHere->",maxEndingHere)
     	}

     if maxSoFar < maxEndingHere {
     	maxSoFar = maxEndingHere
     	fmt.Println("MaxSoFar->",maxSoFar)
     	}
     }
   return maxSoFar
   }

func main() {
	data := []int{1,-2,3,4,-4,6,-14,8,2}
	fmt.Println("Max sub array sum:",MaxSubArraySum(data))
}
