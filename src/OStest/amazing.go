package main

import (
	"fmt"
	"time"
)

func remove(t []int, signal chan int) {
	j := <-signal
	a := 0
	for i := 1; i < len(t); i++ {
		if t[i] == 1 {
			if (i+1)%j == 0 {
				t[i] = 0
				break
			}
			if a == 0 {
				a = i + 1
			}
		}
	}
	fmt.Println(a)
	t[0] = a
}

func main() {
	var n int
	a, err := fmt.Scanf("%d", &n)
	if (a == 1) && (err == nil) {
		fmt.Println(n)
	} else {
		return
	}
	var t []int
	for i := 0; i < n; i++ {
		t = append(t, 1)
	}
	sig := make(chan int)
	go remove(t, sig)
	sig <- 2
	for {
		if t[0] < n {
			go remove(t, sig)
			continue
		}
		break
	}
	time.Sleep(time.Second * 3)
	for i := 0; i < len(t); i++ {
		if t[i] == 1 {
			fmt.Print(i + 1)
		}
	}
}
