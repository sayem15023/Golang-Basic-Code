package main

import "fmt"

type myNumber int

func (num myNumber) square() int {
	if num == 0 {
		return 1
	}
	return int(num * num)
}
func main() {
	num := myNumber(25)
	squ := num.square()
	fmt.Printf("The squere of %d is %d\n", num, squ)
}
