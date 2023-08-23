package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	byt := []byte(
		`
		{
		"num":6.23,
		"strs":["a","b"]
		}
		`,
	)
	var v map[string]interface{}
	if err := json.Unmarshal(byt, &v); err != nil {
		panic(err)
	}
	fmt.Println(v)
}
