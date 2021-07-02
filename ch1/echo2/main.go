// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
)

// 练习1.2	打印索引和值
// 不同于c，若数量级大，如此叠加字符开销较大
func main() {
	s, seq := "", ""
	for i, arg := range os.Args[1:] {
		s += seq + strconv.Itoa(i) + " " + arg
		seq = "\n"
	}
	fmt.Println(s)
}

//!-
