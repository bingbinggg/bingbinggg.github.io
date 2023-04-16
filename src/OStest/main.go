package main

import (
	"fmt"
	"time"
)

func aTask() {
	fmt.Println("a")
}

func Task() {
	fmt.Println("T")
}

func main() {
	go aTask()
	time.Sleep(time.Second * 1)
	Task()
	go aTask()
	time.Sleep(time.Second * 1)
}
