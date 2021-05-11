package main //default package

import "fmt" //priting things on console

func main() { //entry point
	//arrays
	var fruitArr[2] string

	//assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Oranges"

	//declare and assign(shorthand)
	fruitArr := [2]string{"Apple", "Oranges"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	//array slice
	fruitSlice := []string{"Apple", "Oranges", "Grapes", "Cherry"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])

}
