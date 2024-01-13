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
		calculateSquare(i)
	}
	elapsed := time.Since(start)
	fmt.Println("Function took", elapsed)

}