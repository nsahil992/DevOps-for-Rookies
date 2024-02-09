package main
import (
	"fmt"
	"time"
)
func calculateSquare(i int) {
	time.Sleep(1 * time.Second)
	fmt.Println(i*i)
		
}

func main() {

	start := time.Now()
	for i := 1; i <= 10000; i++ {
		go calculateSquare(i)   // go keyword is used when calling a function
		// Here all these 10000 function calls will be printed in 10000 go routines
	}
	elapsed := time.Since(start)
	time.Sleep(2 * time.Second) // our main function is going to wait atleast 2 seconds before it exits
	fmt.Println("Function took", elapsed)
}