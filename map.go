package main //default package

import "fmt" //priting things on console

func main() { //entry point
	
	// //Define map
	// emails := make(map[string] string)

	// //assign key values
	// emails["bob"] = "bob@gmail.com"
	// emails["sharon"] = "sharon@gmail.com"

	//declare map and add key values
	emails := map[string]string{"bob":"bob@gmail.com", "sharon":"sharon@gmail.com"}
	emails ["Mike"] = "mike@gmail.com" 
	fmt.Println(emails) // print the map
	fmt.Println(len(emails)) //print the length of the map
	fmt.Println(emails["bob"]) //print a specific value

	//delete from map
	delete(emails, "bob")
	fmt.Println(emails)
}
