package main

import "fmt"

func main() {
	personSalary := make(map[string]int)
	personSalary["satya"] = 12000
	personSalary["murali"] = 13000
	personSalary["mohan"] = 9000
	fmt.Println("person salary content", personSalary)
	employee := "murali"
	fmt.Println("salary of" , employee ,"is",personSalary[employee])
}
