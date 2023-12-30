package main
import ("fmt")

func main(){
	var name string = "Sahil"
	var city string = "New York"
	var age int = 20

	fmt.Println("Wecome to", city, ", ",name)
	fmt.Printf("Nice to see you in %v\n", city)
	fmt.Printf("You are %d years old\n", age)
	fmt.Printf("Variable age is a type of %T\n", age)

	var x,y string = "Go", "Mascot"   // 2 variable declaration of same type
	fmt.Println(x)
	fmt.Println(y)

	var (
		a string = "Gopher"      // Shorthand way of two different variables
		i int = 8999
	)
	fmt.Println(a)
	fmt.Println(i)

	fruit := "apple"   // short variable declaration
	fmt.Println(fruit)


}