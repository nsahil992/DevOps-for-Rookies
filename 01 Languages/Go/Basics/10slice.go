// Slice is a window on an underlying array
// We can increase the size of slice but not an array
// Every slice has a pointer that indicates the start of the slice
// Length is the number of elements in the slice
// Capacity is the maximum number of elements

package main
import ("fmt")

func main() {
	arr := []int {10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("Slice is: ")
	slice_1 := arr[0:8]
	fmt.Println(slice_1)

	// Slice from other slice

	sub_slice := slice_1[1:3]
	fmt.Println("Sub Slice is: ")
	fmt.Println(sub_slice)

	// Changing value of slice

	slice_1[0] = 11
	fmt.Println("After modification of Slice: ")
	fmt.Println(slice_1)

	// Append 

	slice_2 := []string{"Apple", "Mango", "Watermelon"}
	fmt.Println(slice_2)
	fmt.Println("The appended slice is: ")
	fmt.Println(append(slice_2, "Custard apple", "Papaya"))
	


}