package main

import "fmt"

func main() {

	details := make(map[string]any)

	details["Name"] = "Satya"
	details["Age"] = 28
	details["Surname"] = "Peddireddi"

	//another map

	courses := make(map[string]string)
	courses["CM"] = "Ansible"
	courses["IAC"] = "terraform"
	courses["Orchestrator"] = "k8s"
	courses["Container"] = "Docker"
	courses["Cloud"] = "Azure"

	for key, value := range courses {
		fmt.Println(key, value)
	}
	//print only keys
	delete(courses, "CM")

	for key, _ := range courses {
		fmt.Println(key)
	}

	fmt.Printf("Type of the variable %T:", details)
	fmt.Println("Details: ", details["Name"])
}
