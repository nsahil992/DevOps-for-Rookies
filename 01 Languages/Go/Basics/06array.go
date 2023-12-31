package main
import ("fmt")

func main(){
	var fruits [3]string = [3]string{"Apple", "Mango", "Banana"}
	fmt.Println(fruits)

	marks := [3]int{10, 20, 30}   // shorthand way
	fmt.Println(marks)

	languages := [...]string{"Go", "Swift", "Rust"} //[...] is ellipse which calculates the elements present in the array
	fmt.Println(languages)

	fmt.Println(len(languages))  // Prints the length of languages array
	fmt.Println(languages[0])   // [0] stands for index and will print the name of the language at index 0 i.e Go
	
	marks[2] = 99
	fmt.Println(marks)  // The value of marks array at index 2 i.e 30 will be changed to 99

	// Looping through an array
	var grades [5]int = [5]int{80, 83, 96, 90, 75}
	for i := 0; i < len(grades); i++ {
	fmt.Println(grades[i])
	}

	// Multidimensional array

	arr := [3][2]int{{10, 20}, {30, 40}, {50, 60}}
	// [3] refers to the size of the array
	// [2] refers to each element in a array
	fmt.Println(arr[2][1])

}