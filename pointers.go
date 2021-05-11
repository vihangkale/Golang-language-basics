package main //default package

import "fmt" //printing things on console

func main() { //entry point
	a := 5
	b := &a // assigned the memory address of a
	
	fmt.Println(a,b)
	fmt.Printf("%T\n%T\n", a,b) // prints type of the variable

	//use * to read val from address
	fmt.Println(*b) // points to the value of a and prints it
	fmt.Println(*&a)


	//change val with pointer
	*b = 10
	fmt.Println(a)
}
