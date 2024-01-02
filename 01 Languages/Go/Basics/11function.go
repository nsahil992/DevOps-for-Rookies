package main
import("fmt")

func greeting(a string) {
	fmt.Println("Hey there,",a)
}

// Return type functions

func operations(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}
func main() {
	greeting("Sahil")
	sum, diff := operations(10, 9)
	println(sum," ",diff)
}
