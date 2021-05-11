package main //default package

import ("fmt"    //priting things on console
		"strconv"
		)
//define person struct
// type Person struct {
// 	firstName string
// 	lastName string
// 	city string
// 	gender string
// 	age int
// }

//alternate way to define struct
type Person struct {
firstName, lastName, city, gender string
age int
}


// Greeting method(value receiver)
func(p Person) greet() string {
	return "Hell, my name is "+ p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) //strconv just converts the age to string type
}


//hasBirthday method (pointer receiver)
func(p *Person) hasBirthday() { //age increments everytime their bday comes
	p.age++
}


//getsMarried (pointer receiver)
func(p *Person) getMarried(SpouseLastName string) {
	if p.gender =="m" {
		return
	} else {
		p.lastName = SpouseLastName
	}


}

func main() { //entry point
	
	//init person using struct in descriptive way
	person1 := Person {firstName: "Samantha", lastName: "Smith", city:"Boston", gender:"f", age:25}
	person2 := Person {firstName: "Subir", lastName: "Jadhav", city:"New York", gender:"m", age:21}
	// //ALternative
	// person1 := Person {"Samantha", "Smith", "Boston", "f", 25}


	// fmt.Println(person1)
	// fmt.Println(person1.firstName)
	// person1.age++ //increment age
	// fmt.Println(person1)

	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.getMarried("Kale")

	person2.getMarried("Tommy")

	fmt.Println(person2.greet())
}

