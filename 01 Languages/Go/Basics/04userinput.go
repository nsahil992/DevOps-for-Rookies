package main
import ("fmt")

func main(){
	var user_name string
	var age int
	var like string
	fmt.Println("Enter your name: ")
	fmt.Scan(&user_name)

	fmt.Println("What is your age? ")
	fmt.Scan(&age)

	fmt.Println("What do you like? ")
	fmt.Scan(&like)

	fmt.Println("Hello, ", user_name)
	fmt.Printf("%v is %v years old\n", user_name, age)
	fmt.Printf("%v likes %v\n", user_name, like)

}