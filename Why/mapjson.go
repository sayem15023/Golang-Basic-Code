package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	mapS := map[string]int{"Abir": 29, "kamal": 40}
	mapD, _ := json.Marshal(mapS)
	fmt.Println("The map answer", string(mapD))
}
