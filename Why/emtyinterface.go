package main

import "fmt"

type Mystring string
type Rectangular struct {
	width  float64
	height float64
}

func emptyInterface(i interface{}) {
	fmt.Println("Value given to an empty interface function is of type '%T'with %v\n", i, i)
}
func main() {
	myString := Mystring("My name is something")
	rectangular := Rectangular{4.5, 7.8}
	emptyInterface(myString)
	emptyInterface(rectangular)
}
