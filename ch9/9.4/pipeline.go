/* 练习 9.4: 创建一个流水线程序，支持用channel连接任意数量的goroutine，
在跑爆内存之前，可以创建多少流水线阶段？
一个变量通过整个流水线需要用多久 */

package pipelinepackage

func pipeline(stages int) (in chan int, out chan int) {
	first := make(chan int)
	// 创建 stages 个协程，每个协程传输 testing.B 中的N个数据
	for i := 0; i < stages; i++ {
		in = first
		out = make(chan int)
		go func(in chan int, out chan int) {
			for v := range in {
				out <- v
			}
			close(out)
		}(in, out)
	}
	return first, out
}
