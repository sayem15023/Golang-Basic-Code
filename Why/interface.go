package main

import "fmt"

type Animal interface {
	walk() string
	bark() string
}
type Dog struct {
	w, b string
}

func (d Dog) walk() string {
	return d.w
}
func (d Dog) bark() string {
	return d.b
}
func main() {
	var a Animal = Dog{"Dog is walking", "Dog is barking"}
	fmt.Println("!....Dog...")
	fmt.Println(a.walk())
	fmt.Println(a.bark())
}
