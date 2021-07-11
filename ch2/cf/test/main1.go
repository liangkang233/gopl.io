package main

import (
	"fmt"
	"gopl/ch2/tempconv"
	"os"
	"strconv"
)

func main() {
	// var pc1 [256]byte = func() (pc [256]byte) {
	// 	for i := range pc {
	// 		pc[i] = pc[i/2] + byte(i&1)
	// 	}
	// 	return pc
	// }()
	// fmt.Printf("%v", pc1)

	// 测试练习2.2 导入的英尺与米的转换规则，重量类似不做
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Feet(t)
		m := tempconv.Meter(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToM(f), m, tempconv.MToF(m))
	}
}
