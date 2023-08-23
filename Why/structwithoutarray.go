package main

import "fmt"

type Rool struct {
	x int
	y int
}

//Here we use struct without using array
func main() {
	newStruct := new(Rool)
	newStruct.x = 12
	newStruct.y = 15
	fmt.Println("newstruct", newStruct)
}
