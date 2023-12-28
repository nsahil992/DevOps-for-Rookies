package main
import ("fmt")

func main(){
	var age int
	fmt.Println("Enter your age: ")    // asks user for their age
	fmt.Scan(&age)  // takes age from the user
	if(age >= 18){  // if the age of the user is equal to 18 or greater than 18, it will print the following
		fmt.Println("You are old enough to drive, Welcome!.")
	} else {    // if user age is less than 18, it will print the following
		fmt.Println("Sorry! You are not old enough.")
	}
}