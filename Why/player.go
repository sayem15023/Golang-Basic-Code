package main
import"fmt"
//struct act like a class we can define eveything on it 
//here I create an instance by using struct 
type Player struct{
	name string
	age int         //here name,age,city is known for field or 
	city string
}
func main() {
	var player1 Player
	player1.name = "Tamim"
	player1.city = "ctg"
	player1.age = 29
	fmt.Println(player1)
}
