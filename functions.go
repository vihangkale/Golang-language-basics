package main //default package

import "fmt" //priting things on console

func greet(name string) string {
	return "Hello " + name
}

func getSum(num1 , num2 int) int {
	return num1 + num2
}

func main() { //entry point
	
	fmt.Println(getSum(3, 4)) // syntax to print anything

}
