package main
import "fmt"

func main() {
	x := func(a int, b int)int {
		return a * b
	}
	fmt.Printf("%T\n", x)
	fmt.Println(x(20, 30))
}