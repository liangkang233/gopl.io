// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package popcount

// pc[i] is the population count of i.
// 下面的代码定义了一个PopCount函数，用于返回一个数字中含二进制1bit的个数。它使用init初始化函数来生成辅助表格pc，
// pc表格用于处理每个8bit宽度的数字含二进制的1bit的bit个数，
// 这样的话在处理64bit宽度的数字时就没有必要循环64次，只需要8次查表就可以了。
// （这并不是最快的统计1bit数目的算法，但是它可以方便演示init函数的用法，并且演示了如何预生成辅助表格，这是编程中常用的技术）。

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// 练习2.3 重写PopCount
func MyPopCount(x uint64) int {
	var ans int = 0
	for i := 0; i < 8; i++ {
		ans += int(pc[byte(x>>(i*8))])
	}
	return ans
}

/* 使用匿名函数做init
var pc [256]byte = func() (pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return //返回值显示声明为pc，所以不用函数内声明也可。默认是返回pc变量、return不补齐也可（bare return）  LK233 note
}() */

//!-
