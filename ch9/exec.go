package ch9

import (
	"fmt"
)

func Exec() {
	go teller()
	go func() {
		Disposit(200)
		fmt.Println("bob", Balance())
	}()
	go func() {
		Disposit(100)
		fmt.Println("ali", Balance())
	}()

	fmt.Println("sum", Balance())
}
