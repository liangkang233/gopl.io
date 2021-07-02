// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

//!+ to test func
func useJoin(args []string) string {
	return strings.Join(args[0:], " ")
}

func useAdd(args []string) string {
	s, sep := "", ""
	for _, arg := range args[0:] {
		s += sep + arg
		sep = " "
	}
	return s
}

//!-

//!+
func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(useAdd(os.Args[1:]))
	fmt.Println(useJoin(os.Args[1:]))
}

//!-
