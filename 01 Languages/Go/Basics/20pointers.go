package main
import "fmt"

func main() {
	a := 10
	fmt.Printf("%T %v\n", &a, &a) // to find address of operator
	fmt.Printf("%T %v\n", *(&a), *(&a)) // to dereference operator


	// Declaring a pointer

	var i *int
	var s *string
	fmt.Println(i)
	fmt.Println(s)

	// Initializing a pointer

	b := "sahil"           // shorthand way
	ptr_b := &b
	fmt.Println(ptr_b)

	// Dereferencing and modifying a pointer

	z := "Hello"
	fmt.Printf("%T %v\n", z, z)
	ptr_z := &z
	*ptr_z = "World"
	fmt.Printf("%T %v\n", z, z)	

}