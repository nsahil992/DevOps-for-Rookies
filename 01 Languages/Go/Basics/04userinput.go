package main
import ("fmt")

func main(){
	var a string // String datatype
	var b int   //  Int datatype
	fmt.Println("Enter a string and a number: ")
	count, err := fmt.Scanf("%s %d", &a, &b) // takes user input for a and b
	fmt.Println("count: ", count)           // Checks the count
	fmt.Println("error: ", err)            // Checks the error if user provided invalid datatype 
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

}