package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var boolean = true
	bolB, _ := json.Marshal(boolean)
	for i := range bolB {
		fmt.Print(string(bolB[i]))
	}
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("Ghadha")
	fmt.Println(string(strB))
}
