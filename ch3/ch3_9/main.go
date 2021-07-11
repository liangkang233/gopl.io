/* 练习 3.9： 编写一个web服务器，用于给客户端生成分形的图像。运行客户端用过HTTP参数参数指定x,y和zoom参数
 */

package main

import (
	"fmt"
	"gopl/ch3/server_surface"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/figure", server_surface.Figure_handler) //把输出的gif的wirte操作到返回http请求
	http.HandleFunc("/svg", server_surface.Svg_handler)       //输出的svg的wirte操作到返回http请求
	http.HandleFunc("/man", server_surface.Man_handler)       //输出的mandelbrot的wirte操作到返回http请求

	fmt.Printf("请访问 http://localhost:8000/ 使用加上任意路径返回该get请求内容\n")
	fmt.Printf("添加路径figure和get参数cycles(默认5)可输出随机gif\n")
	fmt.Printf("添加路径svg和get参数index, width, height(默认0,600,320)输出ch3中所需3D图\n")
	fmt.Printf("举例: http://localhost:8000/figure?cycles=1")
	fmt.Printf("举例: http://localhost:8000/svg?index=0&width=600&height=320")
	fmt.Printf("举例: http://localhost:8000/man?xmin=-2&ymin=-2&zoom=1")

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	// 下面程序为获取并打印get或post的form表单(其中get路径的请求也会归纳到form表单)
	// 即http://127.0.0.1:8000/?IO=100&IO=19&ui=789 form表单中会出现
	// Form["IO"] = ["100","19"] Form["ui"] = ["789"]
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
