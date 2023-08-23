package main

import (
	"fmt"
	"math"
)

func main() {
	_, err := math.Sqrt(-1)
	if err != nil {
		if math.IsNaN(err.(math.NaNError).V) {
			fmt.Println("this is not a proper code")
		} else {
			fmt.Println("an error occurred:", err)
		}
	}
}
