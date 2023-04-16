package main

import (
	"fmt"
	"time"
)

func longTask(signal chan int) {
	for {
		fmt.Println("hi")
		v := <-signal
		if v == 1 {
			break
		}
		time.Sleep(time.Second * 1)
	}
	fmt.Println("bye")
}

func main() {
	sig := make(chan int)
	go longTask(sig)
	time.Sleep(time.Second * 10)
	sig <- 1
	time.Sleep(time.Second * 1)

}
