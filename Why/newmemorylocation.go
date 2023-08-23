package main

import "fmt"

func main() {
	//allocated memory for an int using new() function
	p := new(int)
	fmt.Println("value of p is", *p)
	fmt.Println("Memory address of p iss", p)
	*p = 10
	fmt.Println("value of p is ", *p)
	fmt.Println("Memory address of p", p)
}
