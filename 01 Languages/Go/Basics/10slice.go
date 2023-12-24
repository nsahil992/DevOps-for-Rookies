// Slice is a window on an underlying array
// We can increase the size of slice but not an array
// Every slice has a pointer that indicates the start of the slice
// Length is the number of elements in the slice
// Capacity is the maximum number of elements

package main
import ("fmt")

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // will take start(3) and take the elements before end(7)
	fmt.Println(s) // will print 3,5,7

	arr := []string{"a", "b", "c", "d", "e", "f", "g"}
	s1 := arr[1:3]
	s2 := arr[2:5]

	fmt.Println(s1) // will print b,c
	fmt.Println(s2) // will print c,d,e
}