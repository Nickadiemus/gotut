package main

import (
	"fmt"
	"strconv"
)

// define person
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// value reciever method
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.firstName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday pointer reciever method
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	// init person w/ property name
	Person1 := Person{firstName: "Nick", lastName: "Sladic", city: "Dirty L", gender: "M", age: 23}
	fmt.Println(Person1)
	//short hand
	Person2 := Person{"Gatewood", "Robbins", "Dirty L", "M", 22}
	fmt.Println(Person2)

	// accessing properties
	fmt.Println("Happy Birthday " + Person2.firstName)
	// increment age
	Person2.hasBirthday()
	Person2.hasBirthday()
	Person2.hasBirthday()
	fmt.Println(Person2.greet())
}
