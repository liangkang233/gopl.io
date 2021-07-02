// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
// windows下 标准输入不兼容（\n 变为 \r\n），改为直接写入文件
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// 练习1.5 1.6
// var palette = []color.Color{color.White, color.Black}
// var palette = []color.Color{color.White, color.RGBA{0x00, 0xff, 0x00, 0xff}}
var palette = []color.Color{color.White, color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	whiteIndex = iota // first color in palette
	blackIndex        // next color in palette
	greenIndex        // next color in palette
)

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" { //网页访问
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	} else if len(os.Args) > 1 { //生成文件
		f, err := os.Create(os.Args[1] + ".gif")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		lissajous(f)
	} else { // 打印输出
		lissajous(os.Stdout)
	}
}

/*  直接.\lissajous.exe > outfile 编码有误，
*	一种解决方式如上直接生成文件流写入
*	还有一种是写入lissajous字节内存中，再生成到文件
*	buf := &bytes.Buffer{}
*	lissajous(buf)
*	if err := ioutil.WriteFile("a.gif", buf.Bytes(), 0666); err != nil {
*		panic(err)
*	}
 */

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
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
		for t := 0.0; t < cycles*2*math.Pi; t += res {
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

//!-main
