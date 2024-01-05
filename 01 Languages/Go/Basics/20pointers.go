package main
import "fmt"

func main() {
	i := 10
	fmt.Printf("%T %v\n", &i, &i) // to find address of operator
	fmt.Printf("%T %v\n", *(&i), *(&i)) // to dereference operator
}