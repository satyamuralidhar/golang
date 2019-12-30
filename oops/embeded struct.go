package main

import (
	"fmt"
)

type author struct {
	first_name string
	last_name  string
	bio        string
}

func (a author) fullname() string {
	return fmt.Sprintf("%s %s", a.first_name, a.last_name)
	
}
type post struct {
	title string
	content string
	author

}
func (p post) details() {
	fmt.Println("Title: " ,p.title)
	fmt.Println("content:" ,p.content)
	fmt.Println("author:",p.fullname())
}
func main() {
	author1 := author {
		"satya muralidhar",
		"peddireddi",
		"this is my go world",
	}
	post1 := post {
		"satya's golang repo",
		"this is about golang content",
		author1,
	}
	post1.details()
}
