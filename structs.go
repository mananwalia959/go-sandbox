package structs

import (
	"log"
	"time"
)

type User struct {
	firstName string
	lastName  string
	age       int
	birthdate time.Time
}

func main() {
	user := User{
		firstName: "Manan",
		lastName:  "Walia",
		age:       21,
	}
	log.Println(user.firstName)
	log.Println(user.lastName)
	log.Println(user.age)
}
