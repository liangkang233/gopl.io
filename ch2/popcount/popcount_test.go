// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount_test

import (
	"testing"

	// "gopl.io/ch2/popcount"
	"gopl/ch2/popcount"
)

// -- Alternative implementations --

// 没太看懂，不重要
func BitCount(x uint64) int {
	// Hacker's Delight, Figure 5-2.
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}

// 练习2.5	表达式x&(x-1)用于将x的最低一个非零的bit位清零
func PopCountByClearing(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}

// 练习2.4  用移位算法重写PopCount函数，每次测试最右边的1bit，然后统计总数。
func PopCountByShifting(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}
	return n
}

// -- Benchmarks --

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

// goos: windows	BenchmarkPopCount
// goarch: amd64
// pkg: gopl/ch2/popcount
// cpu: Intel(R) Core(TM) i5-10500 CPU @ 3.10GHz
// BenchmarkPopCount-12    	1000000000	         0.2360 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	gopl/ch2/popcount	0.326s

func BenchmarkMyPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.MyPopCount(0x1234567890ABCDEF)
	}
}

// goos: windows	BenchmarkMyPopCount
// goarch: amd64
// pkg: gopl/ch2/popcount
// cpu: Intel(R) Core(TM) i5-10500 CPU @ 3.10GHz
// BenchmarkMyPopCount-12    	219764240	         5.290 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	gopl/ch2/popcount	1.815s

func BenchmarkBitCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCount(0x1234567890ABCDEF)
	}
}

// goos: windows	BenchmarkBitCount
// goarch: amd64
// pkg: gopl/ch2/popcount
// cpu: Intel(R) Core(TM) i5-10500 CPU @ 3.10GHz
// BenchmarkBitCount-12    	1000000000	         0.2513 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	gopl/ch2/popcount	0.356s

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(0x1234567890ABCDEF)
	}
}

// goos: windows	BenchmarkPopCountByClearing
// goarch: amd64
// pkg: gopl/ch2/popcount
// cpu: Intel(R) Core(TM) i5-10500 CPU @ 3.10GHz
// BenchmarkPopCountByClearing-12    	63029512	        19.62 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	gopl/ch2/popcount	1.351s

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0x1234567890ABCDEF)
	}
}

// goos: windows	BenchmarkPopCountByShifting
// goarch: amd64
// pkg: gopl/ch2/popcount
// cpu: Intel(R) Core(TM) i5-10500 CPU @ 3.10GHz
// BenchmarkPopCountByShifting-12    	57874556	        20.53 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	gopl/ch2/popcount	1.290s

// Go 1.6, 2.67GHz Xeon
// $ go test -cpu=4 -bench=. gopl.io/ch2/popcount
// BenchmarkPopCount-4                  200000000         6.30 ns/op
// BenchmarkBitCount-4                  300000000         4.15 ns/op
// BenchmarkPopCountByClearing-4        30000000         45.2 ns/op
// BenchmarkPopCountByShifting-4        10000000        153 ns/op
//
// Go 1.6, 2.5GHz Intel Core i5
// $ go test -cpu=4 -bench=. gopl.io/ch2/popcount
// BenchmarkPopCount-4                  200000000         7.52 ns/op
// BenchmarkBitCount-4                  500000000         3.36 ns/op
// BenchmarkPopCountByClearing-4        50000000         34.3 ns/op
// BenchmarkPopCountByShifting-4        20000000        108 ns/op
//
// Go 1.7, 3.5GHz Xeon
// $ go test -cpu=4 -bench=. gopl.io/ch2/popcount
// BenchmarkPopCount-12                 2000000000        0.28 ns/op
// BenchmarkBitCount-12                 2000000000        0.27 ns/op
// BenchmarkPopCountByClearing-12       100000000        18.5 ns/op
// BenchmarkPopCountByShifting-12       20000000         70.1 ns/op
