package main //default package

import "fmt" //printing things on console

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alias for uint8

// rune // alias for int32 ~= a character (Unicode code point) - very Viking

// float32 float64

// complex64 complex128
// age = 29 //int

//placeholders
// bool:                    %t
// int, int8 etc.:          %d
// uint, uint8 etc.:        %d, %#x if printed with %#v
// float32, complex64, etc: %g
// string:                  %s
// chan:                    %p
// pointer:                 %p


func main() { //entry point
	//using var
	// var name  = "Vings"; //string
	var age int32 = 19 //int
	const isCool = true
	// isCool = false
	var size float32 = 2.5

	//shorthand declaration
	// name:="vings";
	// // size:= 1.5
	// email:="v@gmail.com"

	name, email := "vihang", "v@gmail.com" // shorter way to declare variables

	fmt.Println(name, email,age, isCool, size) //printing the 2 variables
	fmt.Printf("%T\n", isCool) //get the type of varaible, %T gets the type of the variable

}
