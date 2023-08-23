package main

import "fmt"

type Address struct {
	city   string
	street string
	road   string
	zip    string
}
type Customer struct {
	name    string
	age     string
	address Address //for our need we can use a struct as a field
	//if we need multiple instance in a struct then we can use array
}

//if we create an instace of a struct then we can also use the pointer
func main() {
	customer1 := &Customer{
		name: "Safat",
		age:  "27",
		address: Address{
			city:   "London",
			street: "12/ew",
			road:   "34",
			zip:    "2200",
		},
	}
	fmt.Println(customer1)
}
//using pointer we do not need extra space we just located the memory address
