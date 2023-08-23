package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "Sayem", Age: 28}
	file, err := os.Create("person.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(person)
	if err != nil {
		panic(err)
	}
}
