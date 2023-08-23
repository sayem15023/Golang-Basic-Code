package main

import "fmt"

//by using literal struct directly assigned values
type Student struct {
	name string
	city string
	age  int
}

func main() {
	student1 := Student{"Sojib", "Tangail", 22}
	fmt.Println(student1)
}
