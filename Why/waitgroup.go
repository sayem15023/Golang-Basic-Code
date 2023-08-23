package main

import (
	"fmt"
	"sync"
	"time"
)

func routine(waitgroup *sync.WaitGroup, number int) {
	defer waitgroup.Done()
	fmt.Printf("starting routine %d \n", number)
	time.Sleep(time.Second)
	fmt.Printf("Done with routine %d \n", number)
}
func main() {
	fmt.Println("starting main goroutine")
	//create a wait group variable 
	var waitgroup sync.WaitGroup
	for i := 0; i < 5; i++ {
		waitgroup.Add(1)
		go routine(&waitgroup, i)
	}
	waitgroup.Wait()
	fmt.Println("finishing main goroutine")
}
