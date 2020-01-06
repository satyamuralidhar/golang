package main

import "fmt"

type Planet struct {
	Size   int
	Radius int
}

func (p Planet) GetSizeTimeandRadius() int {
	return p.Size * p.Radius
}
func (p *Planet) setSize(size int) {
	p.Size = size
}

func (p *Planet) setRadius(radius int) {
	p.Radius = radius
}
func main() {
	var world Planet = Planet{1, 2}
	world.setSize(5)
	world.setRadius(10)
	fmt.Println(world.GetSizeTimeandRadius())

}
