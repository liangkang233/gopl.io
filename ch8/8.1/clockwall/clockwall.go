/* 练习8.1： 修改clock2来支持传入参数作为端口号，然后写一个clockwall的客户端程序，
这个程序可以同时与多个clock服务器通信，从多个服务器中读取时间，
并且在一个表格中一次显示所有服务器传回的结果，类似于你在某些办公室里看到的时钟墙。
如果你有地理学上分布式的服务器可以用的话，让这些服务器跑在不同的机器上面；
或者在同一台机器上跑多个不同的实例，这些实例监听不同的端口，假装自己在不同的时区。像下面这样：

$ TZ=US/Eastern    ./clock2 -port 8010 &
$ TZ=Asia/Tokyo    ./clock2 -port 8020 &
$ TZ=Europe/London ./clock2 -port 8030 &
$ clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030 */
// clockwall listens to multiple clock servers concurrently.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type clock struct {
	name, host string
}

func (c *clock) watch(w io.Writer, r io.Reader) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		fmt.Fprintf(w, "%s: %s\n", c.name, s.Text())
	}
	fmt.Println(c.name, "done")
	if s.Err() != nil {
		log.Printf("can't read from %s: %s", c.name, s.Err())
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "usage: clockwall NAME=HOST ...")
		os.Exit(1)
	}
	clocks := make([]*clock, 0)
	for _, a := range os.Args[1:] {
		fields := strings.Split(a, "=")
		if len(fields) != 2 {
			fmt.Fprintf(os.Stderr, "bad arg: %s\n", a)
			os.Exit(1)
		}
		clocks = append(clocks, &clock{fields[0], fields[1]})
	}
	for _, c := range clocks {
		conn, err := net.Dial("tcp", c.host)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go c.watch(os.Stdout, conn)
	}
	// Sleep while other goroutines do the work.
	for {
		time.Sleep(time.Minute)
	}
}