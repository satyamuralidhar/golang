/*
structs are nothing but python classes . go dont have concept of classes
no inheritance , super and parent or child concepts.

*/

package main

import "fmt"

func main() {

	type User struct {
		Name   string
		Email  string
		Age    int
		Course string
	}

	details := User{"Satya", "devops.prepare@gmail.com", 28, "DevOps"}

	fmt.Printf("Details: %+v\n", details)
	//+v provides breif details o/p: Details: {Name:Satya Email:devops.prepare@gmail.com Age:28 Course:DevOps}

	fmt.Printf("User Name: %v\n", details.Name)
	fmt.Printf("Email id: %v Course: %v", details.Email, details.Course)
	/*
		o/p:
		User Name: Satya
		Email id: devops.prepare@gmail.com Course: DevOps
	*/

}
