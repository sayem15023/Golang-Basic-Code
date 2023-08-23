package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
}
type MyError struct {
	Code    int
	Message string
}

func main() {
	employee := new(Employee)
	employee.FirstName = "Sayem"
	err := myFunc(*employee)
	fmt.Println(err.Error())

}
func (e MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}
func myFunc(employee Employee) error {
	if employee.FirstName == "" {
		return MyError{Code: 400, Message: "First NAme required"}
	}
	if employee.LastName == "" {
		return MyError{Code: 400, Message: "Last Name Required"}
	}
	return nil
}
