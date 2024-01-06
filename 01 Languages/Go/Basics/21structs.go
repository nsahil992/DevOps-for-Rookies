package main

import "fmt"

type Circle struct {
	x      float64
	y      float64
	radius float64
}

type student struct {
	name   string
	rollNo int
}

type Rectangle struct {
	length  int
	breadth int
}

func main() {
	// for Circle

	var c Circle
	fmt.Printf("%+v\n", c)

	// for struct student
	st := student{
		name:   "Sahil",
		rollNo: 43,
	}
	fmt.Printf("%+v\n", st)

	// Accessing fields of Struct Rectangle

	var r Rectangle
	r.length = 5
	r.breadth = 5
	fmt.Printf("%+v\n", r)
}
