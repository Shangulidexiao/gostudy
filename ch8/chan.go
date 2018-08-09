package ch8

import (
	"fmt"
	"log"
)

func preport() {
	numbers := make(chan int, 3)
	npows := make(chan int, 4)
	go createNumbers(numbers)
	go pows(numbers, npows)
	p(npows)
}

func createNumbers(numChan chan<- int) {
	log.Printf("createNumbers start")
	for i := 1; i < 10; i++ {
		numChan <- i
	}
	close(numChan)
	log.Printf("createNumbers end")
}

func pows(out <-chan int, in chan<- int) {
	log.Printf("pows start")
	for num := range out {
		in <- num * num
	}
	close(in)
	log.Printf("pows end")
}

func p(out <-chan int) {
	log.Printf("p start")
	for num := range out {
		fmt.Println(num)
	}
	log.Printf("p end")
}
