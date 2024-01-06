package main
import "fmt"

type s1 struct {
	x int
}

func main() {
	c := s1{x: 5}
	c1 := s1{x: 5}
	if c == c1 {
		fmt.Println("c is same as c1")
	} else {
		fmt.Println("c is different from c1")
	}
}