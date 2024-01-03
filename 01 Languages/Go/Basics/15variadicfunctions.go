package main
import ("fmt")

func sumNumbers(numbers ...int)int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

	// Eg with more than one parameter

func printDetails(student string, subjects ...string) {
	fmt.Println("Hey", student, ", here are your subjects - ")
	for _, sub := range subjects {
		fmt.Printf("%s, ",sub)
	}
}

func main() {
	fmt.Println(sumNumbers())
	fmt.Println(sumNumbers(10))
	fmt.Println(sumNumbers(10, 20))
	fmt.Println(sumNumbers(10, 20, 30, 40, 50))	


	printDetails("Sahil", "Quantum Physics", "Rocket Science")
}