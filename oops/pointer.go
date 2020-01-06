package main 
import "fmt"
type Planet struct {
	Size int
	Radius int
}

func main() {
	var world *Planet = &Planet{1,2} // inside the planet we have sizes and radius
	mars := world
	mars.Size = 100
	fmt.Println(world.Size)
	fmt.Println(mars.Size)

}

