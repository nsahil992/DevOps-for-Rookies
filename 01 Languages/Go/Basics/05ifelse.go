package main
import ("fmt")

func main(){
	var age int
	fmt.Println("Enter your age: ")
	fmt.Scan(&age)
	if(age >= 18){
		fmt.Println("You are old enough to drive, Welcome!.")
	} else {
		fmt.Println("Sorry! You are not old enough.")
	}


}