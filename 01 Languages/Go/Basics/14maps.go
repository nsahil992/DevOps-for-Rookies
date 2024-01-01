package main
import ("fmt")

func main() {
	language := map[string]string{"en":"English", "fr":"French", "it":"Italian"}
	// (en, fr, it) are the keys and (English, French, Italian) are the values
	fmt.Println(language)

	// Printing length of the map
	fmt.Println(len(language))

	// To print values of the keys
	fmt.Println(language["en"])
	fmt.Println(language["fr"])
	fmt.Println(language["it"])

	// Getting a key
	value, found := language["en"]
	fmt.Println(value, found)       // It will print English and true as English value exists in the map
	value, found = language["hh"]  //  It will print false as hh doesn't exist in the map
	fmt.Println(value, found)

	// Adding key value pair
	language["hi"] = "Hindi"
	fmt.Println(language)

	// Updating key value pair
	language["hi"] = "Hindi language"
	fmt.Println(language)

	// Deleting key value pair
	delete(language, "fr")
	fmt.Println(language)

	// Iterating over a map
	for key, value := range language {
		fmt.Println(key, "=>", value)
	}
}