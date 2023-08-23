package main

import "fmt"

func main() {
	type Vertex struct {
		x int
		y int
	}
	newStruct := new([5]Vertex)
	fmt.Println("length", len(newStruct), "Capacity", cap(newStruct))
	//accessing the 1st element
	newStruct[0].x = 1
	newStruct[0].y = 10
	newStruct[4].x = 5
	newStruct[4].y = 3
	fmt.Println("newstruct", newStruct)
}
