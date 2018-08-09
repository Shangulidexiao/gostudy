package ch8

import (
	"fmt"
	"time"
)

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func action(t time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(t)
		}
	}
}

func ca() {
	go action(100 * time.Microsecond)
	fmt.Println("\r", fib(45))
}
