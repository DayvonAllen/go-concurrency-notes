package channels

import "fmt"

// Implement relaying of message with Channel Direction
func genMsg(ch1 chan<- string) {
	// send message on ch1
	ch1 <- "message"
}

func relayMsg(ch1 <-chan string, ch2 chan<- string) {
	// receive message on ch1
	m := <-ch1
	// send it on ch2
	ch2 <- m
}

func ChannelDirection() {
	// create ch1 and ch2
	ch1 := make(chan string)
	ch2 := make(chan string)

	// spin up goroutine genMsg and relayMsg
	go genMsg(ch1)
	go relayMsg(ch1, ch2)
	// receive message on ch2
	m := <-ch2

	fmt.Println(m)
}
