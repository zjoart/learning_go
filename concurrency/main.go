package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println("Hello, concurrency in Golang")

	var wg sync.WaitGroup

	f := 5
	c := &f

	ch := make(chan int)

	wg.Add(1)

	go sendMessage(f, &wg, ch)

	channelValue := reader(ch)

	wg.Wait()
	close(ch)

	fmt.Println("F & channel values are: ", f, channelValue)
	fmt.Println("F references ", &f, c)
	*c = *c * 2 // do multiplication on the pointer
	fmt.Println("C pointer ", *c)
	fmt.Println("F value now ", f)
	fmt.Print("Done for message to be sent...")
}

// read message
func reader(ch <-chan int) int {

	message := <-ch
	return message
}

// send message
func sender(ch chan<- int, message int) {
	ch <- message
}

func sendMessage(a int, wg *sync.WaitGroup, ch chan int) {
	time.Sleep(2 * time.Second)
	sender(ch, a)
	fmt.Println("Message sent ", a)
	wg.Done()
}
