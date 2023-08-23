package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("Something wrong ")
	if err != nil {
		fmt.Print(err)
	}
}
