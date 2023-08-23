package main

import "fmt"

func main() {
	newArray := new([5]int)
	fmt.Println("length", len(*newArray), "capacity is", cap(*newArray))
	newArray[0] = 10 //assign avalue
	newArray[4] = 20 //assign avalue
	fmt.Println("newarray", newArray)
}
