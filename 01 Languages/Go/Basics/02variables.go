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

}