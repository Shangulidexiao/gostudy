package ch8

import (
	"fmt"
)

//Exec is ch8 main function
func Exec() {
	// tick := time.Tick(1 * time.Second)

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// 	<-tick
	// }
	// fmt.Println("lunch")
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Printf("%d\n", x)
		case ch <- i:
			fmt.Println("c:", i)
		}
	}
}
