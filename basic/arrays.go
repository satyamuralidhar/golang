package main

import (
	"fmt"
)

func main() {

	var courses [5]string
	courses[0] = "DevOps"
	courses[1] = "DevSecOps"
	courses[2] = "SRE"
	courses[3] = "GitOps"
	fmt.Printf("Type of the var %T\n: ", courses)

	//type 2

	var cicd = [3]string{"Jenkins", "Bamboo", "ArgoCd"}
	fmt.Printf("Type : %T\n", cicd)
	fmt.Println(cicd)

	for _, element := range courses {
		fmt.Println(element)
	}

}
