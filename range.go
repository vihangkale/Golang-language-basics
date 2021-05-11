package main //default package

import "fmt" //priting things on console

func main() { //entry point
	
	ids:= []int{343,12,11,312,34}

	//loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//add ids together
	sum:= 0
	for _, id := range ids { //use underscore when i- iterator is not used
		sum +=id
	}
	fmt.Println("Sum", sum)

	//Range with map
	emails := map[string]string{"bob":"bob@gmail.com", "sharon":"sharon@gmail.com"}

	for k, v := range emails { //iterate thru key value 
		fmt.Printf("%s:%s\n", k, v)
	}


	for k := range emails { //iterate thru key
		fmt.Println("Name: " + k)
	}
}


