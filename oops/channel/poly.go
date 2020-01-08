package main

import (
	"fmt"
)

//interface
type employee interface {
	develop() int
	name() string
}

//structure
type team1 struct {
	totalapp_1 int
	name_1     string
}

func (t1 team1) develop() int {
	return t1.totalapp_1
}
func (t1 team1) name() string {
	return t1.name_1
}

type team2 struct {
	totalapp_2 int
	name_2     string
}

func (t2 team2) develop() int {
	return t2.totalapp_2
}
func (t2 team2) name() string {
	return t2.name_2
}
func finaldevelop(i []employee) {

	totalproject := 0
	for _, elk := range i {

		fmt.Printf("\nProject environment = %s\n ", elk.name())
		fmt.Printf("Total number of project %d\n ", elk.develop())
		totalproject += elk.develop()
	}
	fmt.Printf("\nTotal projects completed by "+
		"the company = %d", totalproject)
}
func main() {

	res1 := team1{totalapp_1: 20,
		name_1: "Android"}

	res2 := team2{totalapp_2: 35,
		name_2: "IOS"}

	final := []employee{res1, res2}
	finaldevelop(final)

}
