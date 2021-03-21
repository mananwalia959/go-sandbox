package channel

import (
	"log"
)

func fibonacci(n int, fibbonaciChannel chan int) {
	defer close(fibbonaciChannel)
	first := 0
	second := 1
	for i := 0; i < n; i++ {
		fibbonaciChannel <- first
		newVal := first + second
		first = second
		second = newVal
	}
}

func main() {

	fibbonaciChannel := make(chan int)
	go fibonacci(20, fibbonaciChannel)

	for num := range fibbonaciChannel {
		log.Println(num)
	}
}
