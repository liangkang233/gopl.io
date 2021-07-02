// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

/* 这个程序从两个package中导入了函数，net/http和io/ioutil包，
http.Get函数是创建HTTP请求的函数，如果获取过程没有出错，那么会在resp这个结构体中得到访问的请求结果。
resp的Body字段包括一个可读的服务器响应流。ioutil.ReadAll函数从response中读取到全部内容；将其结果保存在变量b中。
resp.Body.Close关闭resp的Body流，防止资源泄露，Printf函数会将结果b写出到标准输出流中。 */

func main() {
	for _, url := range os.Args[1:] {
		// 练习1.8
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		// if !strings.HasSuffix(url, ".com") {
		// 	url += ".com"
		// }
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// 练习1.7
		// b, err := ioutil.ReadAll(resp.Body)
		len, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		// 练习1.9
		fmt.Printf("get return status: %s\nbody length:	%d\n", resp.Status, len)
		resp.Body.Close()
		// fmt.Printf("%s", b)
	}
}

//!-
