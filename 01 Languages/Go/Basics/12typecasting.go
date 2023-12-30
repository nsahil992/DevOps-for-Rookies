// Type casting is the process of converting one data type to another.

package main
import "fmt"

func main() {
	// integer to float
	var a int = 90
	var b float64 = float64(a)
	fmt.Printf("%.2f\n", b)

	//float to integer
	var c float64 = 49.50  // can convert float to int but doesn't provide accuracy
	var d int = int(c)
	fmt.Printf("%v\n", d)
}