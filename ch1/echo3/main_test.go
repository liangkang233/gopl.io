/* 练习1.3： 做实验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异。
（1.6节讲解了部分time包，11.4节展示了如何写标准测试程序，以得到系统性的性能评测。） */

package main

import "testing"

//!+bench
func BenchmarkUseAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useAdd([]string{"hello", "my", "name", "is", "lk233"})
	}
}

// result BenchmarkUseAdd
// goos: windows
// goarch: amd64
// pkg: gopl/ch1/echo3
// cpu: Intel(R) Core(TM) i5-10500 CPU @ 3.10GHz
// BenchmarkUseAdd-12    	 6729266	       176.3 ns/op	      64 B/op	       4 allocs/op
// PASS
// ok  	gopl/ch1/echo3	1.455s

func BenchmarkUseJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useJoin([]string{"hello", "my", "name", "is", "lk233"})
	}
}

// result BenchmarkUseJoin
// goos: windows
// goarch: amd64
// pkg: gopl/ch1/echo3
// cpu: Intel(R) Core(TM) i5-10500 CPU @ 3.10GHz
// BenchmarkUseJoin-12    	19365770	        62.02 ns/op	      24 B/op	       1 allocs/op
// PASS
// ok  	gopl/ch1/echo3	1.329s

//!-bench
