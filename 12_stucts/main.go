package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// GREETING METHOD (VALUE RECIEVER)

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and i am " + strconv.Itoa(p.age) //STRING CONV
}

//has brithday method (pointer reciever)

func (p *Person) hasBirthday() {
	p.age++

}

func main() {
	//init person using struct

	//ALTERNATIVE
	person1 := Person{"austin", "mcclarren", "orlando", "M", 25}

	// person1 := Person{
	// 	firstName: "austin", lastName: "mcclarren",city: "orlando", gender: "M", age: 25,
	// }

	// fmt.Println(person1)

	person1.hasBirthday()
	fmt.Println(person1.greet())
}
