package main

import "fmt"

func main() {
	var courses = []string{
		"DevOps",
		"Cloud",
		"Monitoring"}
	fmt.Printf("Tyep : %T\n", courses)

	courses = append(courses, "InfluxDB", "Telegraf")
	fmt.Println(courses)

	courses = append(courses[2:])
	fmt.Println(courses)
	fmt.Println(courses[1:])
	/*
		outputs:
		================
		[DevOps Cloud Monitoring InfluxDB Telegraf]
		[Monitoring InfluxDB Telegraf]
		[InfluxDB Telegraf]

	*/
	//create a slice using make

	//var score = make([]int,4) or
	score := make([]int, 5)

	score[0] = 000
	score[1] = 111
	score[2] = 222
	score[3] = 333
	score[4] = 444

	fmt.Println(score)
	/*notes: we assgined memory with 5 its will store only 5 values but
	by using append will arrange another memory with 5 .

	output:
	==========

	[0 111 222 333 444]
	[0 111 222 333 444 555 666 777 888 999

	*/

	score = append(score, 555, 666, 777, 888, 999)

	fmt.Println(score)

}
