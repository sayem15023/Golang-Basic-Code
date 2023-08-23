package main

import (
	"fmt"
	"time"
)

func worker(cancel <-chan struct{}) {
	fmt.Println("worker started")
	defer fmt.Println("worker stopped")
	for {
		select {
		case <-cancel:
			fmt.Println("Worker canceled")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("Working")
		}
	}
}
func main() {
	cancel := make(chan struct{})
	go worker(cancel)
	time.Sleep(5 * time.Second)
	fmt.Println("Canciling worker")
	close(cancel)
	time.Sleep(2 * time.Second)
	fmt.Println("now main goroutine stopped")

}
