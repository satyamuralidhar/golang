package main

import (
	"fmt"
)

//var name string = "Satyamuralidhar"
/*
Structs are nothing but classes in python
interfaces are polymorphism one user has multiple behaviuors
methods are same as python methods -> functions which are called inside calls we can say its has methods
we can use methods inside of Structs
*/

/*
Structs:
-----------

type Name struct {
	name   string
	age    int
	domain string
}

func main() {
	naam := Name{"satyamuralidhar", 29, "DevOps"}
	fmt.Printf("Hi This is %v\n", naam.name)
	fmt.Printf("Age is: %v\n", naam.age)
	fmt.Printf("Domain is: %v\n", naam.domain)
}
*/

// interface
type Volumeinter interface {
	Volume() float64
}

//creating a class for Cone and Cylinder

type Cylinder struct {
	hight, radius float64
}

type Cone struct {
	hight, radius float64
}

//Creating a method for above methods from the interface and add it into reciver type.

//result := None

func (cy Cylinder) Volume() float64 {
	return ((0.33333333333) * (3.14) * (cy.radius * cy.radius) * cy.hight)

}

func (co Cone) Volume() float64 {
	return ((3.14) * (co.radius * co.radius) * co.hight)
}

//create a funtion to calling the interface

func callInter(vi Volumeinter) {
	fmt.Println(vi.Volume())

}

//calling main funtion

func main() {
	cylinder := Cylinder{hight: 20, radius: 30}
	callInter(cylinder)
	cone := Cone{hight: 20, radius: 30}
	callInter(cone)

	//both will get same o/p
	// cylinder := Cylinder{20, 30}
	// callInter(cylinder)
	// cone := Cone{20, 30}
	// callInter(cone)

}
