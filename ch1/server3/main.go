// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/figure", figure_handler) //把输出的gif的wirte操作到返回http请求
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!+handler
// handler echoes the HTTP request.
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

func figure_handler(w http.ResponseWriter, r *http.Request) {
	cycles := 5
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		if k == "cycles" {
			var err error
			cycles, err = strconv.Atoi(v[0]) //get发送的同名变量cycles只接受第一个
			if err != nil {
				fmt.Fprintf(os.Stderr, "service3 str conv form: %v\n", err)
				os.Exit(1)
			}
			break
		}
	}
	lissajous(w, cycles)
}

// 练习1.12
// 为了不干扰原来的头部form表单的数据返回，添加路由figure
// 修改lissajous,使得访问类似 http://localhost:8000/figure?cycles=20时，会将cycles默认的5改为20
var palette = []color.Color{color.White, color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	whiteIndex = iota // first color in palette
	blackIndex        // next color in palette
	greenIndex        // next color in palette
)

func lissajous(out io.Writer, myCycles int) {
	cycles := float64(myCycles) //直接用cycles int变量，在之后会把math.pi给截断为int
	const (
		// cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < math.Pi*cycles*2; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-handler
