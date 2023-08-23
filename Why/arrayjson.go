package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	str := []string{"Abir", "Sayem", "Rifat"}
	fmt.Println("before json", str)
	str1, _ := json.Marshal(str)
	fmt.Println("After Json", string(str1))

}
