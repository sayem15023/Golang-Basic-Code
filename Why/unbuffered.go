package main

import (
	"fmt"
	"time"
)

// send data into channel
func sendData(ch chan string) {
	fmt.Println("Sending data into channel")
	ch <- "Hello"
	fmt.Println("Strring has been received from channel...")
}
//receive data into channel
func getData(ch chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println("String received from the channel", <-ch)
}
func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	fmt.Scanln()
}
