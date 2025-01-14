// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 244.

// Countdown implements the countdown for a rocket launch.
package main

import (
	"fmt"
	"time"
)

// tick每隔一秒会把当前的time传入ch

//!+
func main() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		// <-tick
		fmt.Println(<-tick)
	}
	launch()
}

//!-

func launch() {
	fmt.Println("Lift off!")
}
