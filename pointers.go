package pointers

import "log"

func main() {
	message := "blue"

	log.Println(" old value is ", message)
	log.Println(&message)
	log.Println(*(&message))
	changeValueByReference(&message)
	log.Println(&message)
	log.Println(*(&message))
	log.Println(" new value is ", message)
}

func changeValueByReference(s *string) {

	newVal := "Red"
	*s = newVal
}
