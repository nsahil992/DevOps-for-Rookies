package main
import ("fmt")

func main(){
	var i int = 1000000000
	var f float64 = 8.91
	var s string = "Sahil"
	var b bool = true

	// assigning value afterwards

	var num int
	num = 2

	// shorthand
	// type 1
	a := "Argo"
	fmt.Println(a)

	//type 2
	var y,z int = 98, 1
	fmt.Println(y, z)

	//type 3
	var (
		 fruit string = "Apple"
		 price int = 100 
	)
	fmt.Println(fruit, price)

	fmt.Println(num)
	fmt.Printf("Hello %v, your CGPA is %v.\n", s, f)
	fmt.Println(b)
	fmt.Println(i)




}