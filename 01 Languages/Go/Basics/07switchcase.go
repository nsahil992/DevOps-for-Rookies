package main
import ("fmt")

func main() {

	var fruits string
	fmt.Println("Enter a name of fruit: ")
	fmt.Scan(&fruits) // takes a name of the fruit from the user
	switch fruits {
	case "Apple": // if user inputs apple, it will print a red fruit
		fmt.Println("A red fruit")
	case "Mango": 
		fmt.Println("King of fruits")
	case "Strawberry":
		fmt.Println("Good for teeth")
	case "Banana":
		fmt.Println("Good for energy")
	default: // if the user enters the name of the fruit which is not mentioned in the case then it will print the default statement
		fmt.Println("Sorry, I didn't found the fruit.")
	}
}