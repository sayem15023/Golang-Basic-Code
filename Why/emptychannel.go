package main

import "fmt"

func main() {
	var ch chan int
	if ch == nil {
		fmt.Println("channel ch is nil")
	}
	ch = make(chan int)
	fmt.Printf("type of %T", ch)
}
