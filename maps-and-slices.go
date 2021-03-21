package mapsandslices

import "log"

type Student struct {
	id   int
	name string
}

func main() {

	names := []string{"Manan", "Aman", "Anam", "Naam", "Maan"}

	student1 := Student{1, "Manan"}
	student2 := Student{2, "Aman"}

	students := []Student{student1, student2}

	idToStudent := make(map[int]Student)

	for _, student := range students {
		idToStudent[student.id] = student
	}

	log.Println(names[2:3])
	log.Println(students)
	log.Println(idToStudent)
}
