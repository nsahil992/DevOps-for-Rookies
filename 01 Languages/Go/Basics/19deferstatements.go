package main
import "fmt"

func printName(a string){
	fmt.Println(a)
}

func printRollno(rno int) {
	fmt.Println(rno)
}

func printAddress(add string) {
	fmt.Println(add)
}

func main() {
	printName("Sahil")
	defer printRollno(43)
	printAddress("Boat Club road")
}