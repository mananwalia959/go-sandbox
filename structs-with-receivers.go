package receivers

import "log"

type User struct {
	name string
}

func (u *User) printName() {
	log.Println(u.name)
}

func main() {
	user1 := User{name: "Mary Manager"}
	user1.printName()

}
