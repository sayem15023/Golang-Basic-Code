package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func (u User) userDetail() string {
	return fmt.Sprintf("The Name is:%s\nEmail is:%s", u.Name, u.Email)
}
func (u *User) changeDetails(newName, newEmail string) {
	u.Name = newName
	u.Email = newEmail
}
func main() {
	user1 := User{Name: "Sayem Chowdhury", Email: "chowdhurysayem23@gmail.com"}
	fmt.Println(user1.userDetail())
	user1.changeDetails("Sakib", "sakib75@gmail.com")
	fmt.Println(user1.userDetail())
}
