package main
import ("fmt"
)
func main() {
	var personSalary map[string]int
	if personSalary == nil {
		fmt.Println("nil value")
		personSalary = make(map[string]int)
	}
}
