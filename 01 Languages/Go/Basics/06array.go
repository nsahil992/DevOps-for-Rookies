package main
import ("fmt")

func main(){
	var Languageslist [4]string // array
	
	Languageslist[0] = "Go"
	Languageslist[1] = "Rust"
	Languageslist[2] = "Swift"
	Languageslist[3] = "Julia"

	fmt.Println(Languageslist) // prints array with 4 languages
	fmt.Println(len(Languageslist)) // prints the length of array i.e 4

	// append inserts names in array
	names := []string{"Sahil", "Courage"} //empty string
	fmt.Println(names)
	fmt.Println(len(names))
	names = append(names, "Phineas") // Phineas is inserted in array
	fmt.Println(names)
	fmt.Println(len(names)) // prints the new length of the 'names' array.
}