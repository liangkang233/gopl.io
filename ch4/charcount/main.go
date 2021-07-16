// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	kinds := make(map[string]int) // 统计unicode种类

	// 练习4.8 使用unicode.IsLetter等相关的函数，统计字母、数字等Unicode中不同的字符类别。
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		for kind, rangeTable := range unicode.Properties {
			// if unicode.In(r, rangeTable) {
			// 	kinds[kind]++
			// }
			if unicode.Is(rangeTable, r) {
				kinds[kind]++
			}
		}
		// unicode.IsLetter()
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
	for kind, n := range kinds {
		fmt.Printf("%-30s %d\n", kind, n)
	}
}

//!-

/* 测试结果
PS E:\Task\Study_Go\gopl.io\ch4> go run .\charcount\main.go
你好 梁康233
test count & kind
家乡 の sakura又开了
heihei
^Z
rune    count
'a'     2
'开'    1
'\n'    4
'n'     2
'乡'    1
'&'     1
'了'    1
'你'    1
'3'     2
'o'     1
'i'     3
'の'    1
's'     2
'k'     2
'd'     1
't'     3
'又'    1
'好'    1
'梁'    1
'u'     2
'h'     2
' '     6
'康'    1
'\r'    4
'家'    1
'r'     1
'2'     1
'e'     3
'c'     1

len     count
1       43
2       0
3       10
4       0
Unified_Ideograph              9
White_Space                    14
Pattern_White_Space            14
Hex_Digit                      10
ASCII_Hex_Digit                10
Pattern_Syntax                 1
Soft_Dotted                    3
Ideographic                    9
*/
