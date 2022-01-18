/* 练习 9.5: 写一个有两个goroutine的程序，两个goroutine会向两个无buffer
channel反复地发送ping-pong消息。
这样的程序每秒可以支持多少次通信？ */

package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	q := make(chan int)
	var i int64
	start := time.Now()
	go func() {
		q <- 1
		for {
			i++
			q <- <-q
		}
	}()
	go func() {
		for {
			q <- <-q
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println(float64(i)/float64(time.Since(start))*1e9, "round trips per second")
}
