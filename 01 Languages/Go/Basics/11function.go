package main
import("fmt")
func main() {
	greetuser() // we called the function
}


func greetuser(){ // greetuser is the name of the function
	fmt.Println("Enter your name: ")
	var username string
	fmt.Scan(&username)
	fmt.Println("Greetings for the day", username)
}

