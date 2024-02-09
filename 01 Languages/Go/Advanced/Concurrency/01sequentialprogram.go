package main
import (
	"fmt"
	"time"
)
func calculateSquare(i int) {
	time.Sleep(1 * time.Second) // timeout that will sleep for 1 sec and then print the result
	fmt.Println(i*i) // function where we will print square of the numbers
		
}

func main() {

	start := time.Now() // will return current local time
	for i := 1; i <= 10000; i++ { // we will print the numbers from 1 to 10000
		calculateSquare(i) // we called the function
	}
	elapsed := time.Since(start) // since is the method from the time package that would return 
	// the time that has elapsed from the current time
	fmt.Println("Function took", elapsed)

}