package main
import ("fmt"
	"strconv")

func main() {
	// conversion integer to string
	var i int = 32
	var s string = strconv.Itoa(i)
	fmt.Printf("%q\n", s)

	// conversion of string to integer
	var a string = "200"
	b, err := strconv.Atoi(a)
	fmt.Printf("%v\n", b, err)

}