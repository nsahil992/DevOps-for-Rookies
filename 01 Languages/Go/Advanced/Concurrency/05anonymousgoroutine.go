package main
import ("fmt"
		"time")

func main() {
	go func() {
		fmt.Println("It's anonymous")
	}()
	time.Sleep(1 * time.Second)
}