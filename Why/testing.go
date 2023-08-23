package main

import ( "testing"
	"fmt")

func DivisibleByThree(num int) string {
	if num%3 == 0 {
		return "Hurry"
	}
	return "Not Hurry"
}

func TestDivisibleByThree(t *testing.T) {
	got := DivisibleByThree(14)
	want := "Hurry"
	if want != got {
		t.Errorf("Expected '%s', but got '%s'", want, got)
	}
}
func main() {
	num := 14
	result := DivisibleByThree(num)
	fmt.Printf("%d is %s\n", num, result)
}
