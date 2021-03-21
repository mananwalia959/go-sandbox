package interfaces

import "log"

type Animal interface {
	says() string
	legs() int
}

type Dog struct {
	name  string
	breed string
}

type Cat struct {
	name  string
	color string
}

func describeAnimal(animal Animal) {
	log.Println("the animal says ", animal.says(), " and it has ", animal.legs(), " legs")
}

func main() {
	dog := Dog{"Tommy", "Labradoe"}
	describeAnimal(dog)
	describeAnimal(Cat{"Mrs Lawrence", "White"})
}

func (dog Dog) says() string { return "woof" }
func (cat Cat) says() string { return "meow" }
func (dog Dog) legs() int    { return 4 }
func (cat Cat) legs() int    { return 4 }
