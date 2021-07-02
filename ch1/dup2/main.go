// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

// 练习1.4 修改dup2 出现重复行时打印文件名称
type countNum struct {
	nums int //行出现的次数
	// filename []string //出现书名的string切片
	filename map[string]int //出现书名的map，自动去重
}

func main() {
	counts := make(map[string]countNum)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.nums > 1 {
			fmt.Printf("\"%d %s\"find \non files:%v\n", n.nums, line, n.filename)
		}
	}
}

func countLines(f *os.File, counts map[string]countNum) {
	input := bufio.NewScanner(f)
	for input.Scan() { //将文件流以行做数据读入map
		_, ok := counts[input.Text()]
		if ok {
			temp := counts[input.Text()]
			temp.nums++
			// temp.filename = append(temp.filename, f.Name())
			temp.filename[f.Name()]++
			counts[input.Text()] = temp
		} else {
			// tmp := countNum{0, nil}
			tmp := countNum{0, make(map[string]int, 0)}
			counts[input.Text()] = tmp
		}

	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
