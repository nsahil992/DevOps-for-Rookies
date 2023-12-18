package main
import ("fmt")

func main(){
	var num int
	fmt.Println("Enter a number to count upto: ")	
	fmt.Scan(&num) // we take number to count upto from the user
	for i := 0; i<=num; i++{ // start num is 0, print the numbers till i = the number given by the user to count upto and we will increase i by 1 everytime
		fmt.Println(i)
	}
}