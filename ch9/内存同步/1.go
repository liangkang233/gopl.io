package main

import (
	"fmt"
)

func main() {
	var x, y int
	ch := make(chan struct{}, 0)
	go func() {
		x = 1                   // A1
		fmt.Print("y:", y, " ") // A2
		ch <- struct{}{}
	}()
	go func() {
		y = 1                   // B1
		fmt.Print("x:", x, " ") // B2
		ch <- struct{}{}
	}()
	<-ch
	<-ch
	close(ch)
}
