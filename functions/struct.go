package main
import "fmt"
type Student struct {
	name string
	rollno int
	percentage float64
}
func main() {
	var satya = Student{"satya muralidhar",458504800,72.6}
	fmt.Println(satya)
	fmt.Println(satya.name)
	fmt.Println(satya.rollno)
	fmt.Println(satya.percentage)
}
