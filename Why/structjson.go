package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type response1 struct {
		Page   int
		Fruits []string
	}
	res1D := &response1{
		Page:   1,
		Fruits: []string{"Apple", "Pear"},
	}
	resB, _ := json.Marshal(res1D)
	fmt.Println(string(resB))
}
