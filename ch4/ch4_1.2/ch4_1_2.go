package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	test4_2()
	if len(os.Args) == 2 {
		fmt.Println("输入两个参数进行哈希对比")
		c1 := sha256.Sum256([]byte(os.Args[1]))
		c2 := sha256.Sum256([]byte(os.Args[2]))
		fmt.Printf("%x\n%x\nc1==c2:%t\n", c1, c2, c1 == c2)
		fmt.Printf("哈希对比位数不同bit位数为%v", test4_1(c1, c2))
	}
}

// 利用匿名函数建8位数的查询表
var pc [256]byte = func() (temp [256]byte) {
	for i := range temp {
		temp[i] = temp[i/2] + byte(i&1)
	}
	return temp
}()

// 练习4.1 计算两个SHA256哈希码中不同bit的数目。（参考2.6.2节的PopCount函数。)
func test4_1(hash1, hash2 [32]byte) int {
	count := 0
	for i := 0; i < len(hash1); i++ {
		count += int(pc[hash1[i]^hash2[i]])
	}
	return count
}

var width = flag.Int("w", 256, "hash width (256 or 512)")

// 练习4.2 打印打印标准输入的SHA256编码，并支持通过命令行flag定制，输出SHA384或SHA512哈希算法。
// stdin输入 flag使用	LK233note
func test4_2() {
	flag.Parse()
	var function func(b []byte) []byte
	switch *width {
	case 256:
		function = func(b []byte) []byte {
			h := sha256.Sum256(b)
			return h[:]
		}
	case 512:
		function = func(b []byte) []byte {
			h := sha512.Sum512(b)
			return h[:]
		}
	default:
		log.Fatal("Unexpected width specified.")
	}
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", function(b))
}
