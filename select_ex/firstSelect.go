package select_ex

import (
	"fmt"
	"time"
)

func FirstSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()

	// multiplex recv on channel - ch1, ch2
	for i:=0; i < 2; i++{
		select {
		case m1 := <-ch1:
			fmt.Println("Received on channel 1: ", m1)
		case m2 := <-ch2:
			fmt.Println("Received on channel 2: ", m2)
		}
	}
}
