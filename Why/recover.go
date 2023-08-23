package main

import "fmt"

func main() {
	accessSlice([]int{1, 2, 3, 4, 5, 6}, 0)
}
func accessSlice(slice []int, index int) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("internal error: %v", p)
		}
	}()
	fmt.Printf("item %d ,value %d \n", index, slice[index])
	defer fmt.Printf("\n defer %d", index)
	accessSlice(slice, index+1)
}
