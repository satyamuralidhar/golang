package main
import ("fmt"
		"time")
func satya() {
	for i := 1 ; i <10 ; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%d",i)
	}
}	
func murali() {
	for i := 'a' ; i < 'g' ; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c",i)
	}
}	
func main() {
	go satya()
	go murali()
	time.Sleep(3500 * time.Millisecond)
	fmt.Println("\n")
	fmt.Println("Hi satya Muralidhar")
}

