package main
import ("fmt"
)
func countries() []string {
	countries := []string{"india","china","america","canada","italy"}
	countryslice := countries[:len(countries)-2]
	countrycopy :=make([]string,len(countryslice))
	copy(countrycopy , countryslice)
	return countrycopy
}
func main() {
countryneed := countries()
fmt.Println(countryneed)
