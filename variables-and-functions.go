package variables

import "log"

func main() {
	var greeting string
	var noun string
	var message string

	greeting, noun = getString("Hello")

	log.Println(greeting + " " + noun)

	message, _ = getString("Goodbye")

	log.Println(message)
}

func getString(s string) (string, string) {
	return s, "World"
}
