// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		// fmt.Printf("  %s\n", comma(os.Args[i]))
		// fmt.Printf("  %s\n", comma1(os.Args[i]))
		fmt.Printf("  %s\n", comma2(os.Args[i]))
	}
	if len(os.Args) == 3 {
		fmt.Printf("%v\n", comma3(os.Args[1], os.Args[2]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
// 函数的功能是将一个表示整数值的字符串，每隔三个字符插入一个逗号分隔符，
// 例如“12345”处理后成为“12,345”。
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

/* LK233 note
bytes包也提供了很多类似功能的函数，[]byte类型和字符串有着相同结构的。
字符串是只读的，逐步构建字符串会导致很多分配和复制。
在这种情况下，使用bytes.Buffer类型将会更有效， */

// 练习3.10 使用迭代代替递归，使用bytes.Buffer代替字符串链接操作。
func comma1(s string) string {
	var buffer bytes.Buffer
	n, t := len(s), len(s)%3
	for i := 1; i <= n; i++ {
		buffer.WriteByte(s[i-1])
		if i%3 == t && i != n {
			buffer.WriteByte(',')
		}
	}
	return buffer.String()
}

// 练习3.11 支持浮点数和可选的正负号处理
func comma2(s string) string {
	if s == "" {
		return s
	}
	var buffer bytes.Buffer
	if s[0] == '+' || s[0] == '-' {
		buffer.WriteByte(s[0])
		s = s[1:]
	}
	arr := strings.Split(s, ".")
	buffer.WriteString(comma1(arr[0]))
	if len(arr) > 1 {
		buffer.WriteByte('.')
		buffer.WriteString(comma1(arr[1]))
	}
	return buffer.String()
}

// 练习3.12 判断两个字符串是否打乱
func comma3(s string, s1 string) bool {
	//LK233 note
	// 由于go对utf8编码的支持，rune相当于每个字符(自动查找该字符有几位，字符串的range也是如此)，而byte一直都是一个字节
	myset := make(map[rune]int, 0)
	myset1 := make(map[rune]int, 0)

	for _, t := range s {
		myset[t]++
	}
	for _, t := range s1 {
		myset1[t]++
	}
	if len(myset) != len(myset1) {
		return false
	}
	for i, j := range myset {
		if myset1[i] != j {
			return false
		}
	}
	return true
}

//!-
