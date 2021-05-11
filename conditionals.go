package main //default package

import "fmt" //priting things on console


// %T	a Go-syntax representation of the type of the value

//placeholders
// bool:                    %t
// int, int8 etc.:          %d
// uint, uint8 etc.:        %d, %#x if printed with %#v
// float32, complex64, etc: %g
// string:                  %s
// chan:                    %p
// pointer:                 %p


func main() { //entry point
	
		
x:= 3
y:= 4

//if else
if x < y || x == y{

	fmt.Printf("%d is less than %d\n", x, y); //placeholders
} else {
	fmt.Printf("%d is less than %d\n", y, x); //placeholders
}


//else if
color:= "red"

if color == "red" {
	fmt.Println("color is red")
} else if color == "blue" {
		fmt.Println("color is blue")

} else {
		fmt.Println("color is not blue nor red")
}




//Switch
switch color {
case "red":
		fmt.Println("color is red")
case "blue":
			fmt.Println("color is red")

default:
		fmt.Println("colorless")

}

}