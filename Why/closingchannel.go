package main

import "fmt"

func producer(cha1 chan int) {
	for i := 0; i < 10; i++ {
		cha1 <- i
	}
	close(cha1)
}
func main() {
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}
}
