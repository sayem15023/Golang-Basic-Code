package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Object interface {
	Volume() float64
}
type Cylinder struct {
	radius float64
	height float64
}

func (c Cylinder) Area() float64 {
	return 2 * math.Pi * c.radius * (c.radius + c.height)
}
func (c Cylinder) Volume() float64 {
	return math.Pi * c.radius * c.radius * c.height
}
func main() {
	var shape Shape = Cylinder{8.5, 4.5}
	cylinder := shape.(Cylinder)
	fmt.Println("Area of a shape of interface type ", cylinder.Area())
	fmt.Println("Area of a object of interface type ", cylinder.Volume())

}
