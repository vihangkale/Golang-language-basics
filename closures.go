package main //default package

import "fmt" //priting things on console

func adder() func(int) int {
	sum:= 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() { //entry point
	sum := adder()

	for i:= 0; i<10; i++ {
		fmt.Println(sum(i))
	}

}
