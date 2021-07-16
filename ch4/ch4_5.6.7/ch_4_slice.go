package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	// Interactive test of reverse.
	input := bufio.NewScanner(os.Stdin)
	// outer:
	for input.Scan() {
		test := []byte(input.Text())
		test1 := []byte(input.Text())
		fmt.Printf("shanjian %s\n", shanjian(test))
		rev(test1)
		fmt.Printf("rev %s\n", test1)
		s := strings.Fields(input.Text())
		fmt.Printf("chuchong%v\n", chuchong(s))
		// NOTE: ignoring potential errors from input.Err()
	}
}

// 练习4.5： 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
func chuchong(s []string) []string {
	i := 0
	myset := make(map[string]int, 0)
	for _, str := range s {
		if myset[str] == 0 {
			myset[str] = 1
			s[i] = str
			i++
		}
		// 使用append方法
		// s = append(s, str)   这样就不是共享一个内存空间了
	}
	return s[:i]
}

// 练习4.6： 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格,替换成一个空格返回
// 参考unicode.IsSpace()
func shanjian(s []byte) []byte {
	i, flag := 0, false
	for _, str := range s {
		if unicode.IsSpace(rune(str)) && flag {
			continue
		} else if unicode.IsSpace(rune(str)) {
			flag = true
		} else {
			flag = false
		}
		s[i] = str
		i++
	}
	return s[:i]
}

// 练习4.7： 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存
func rev(b []byte) {
	size := len(b)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[size-1-i] = b[size-1-i], b[i]
	}
}
