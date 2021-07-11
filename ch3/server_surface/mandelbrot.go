package server_surface

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"net/http"
	"strconv"
)

const (
	m_width, m_height = 1024, 1024
)

var params = map[string]float64{
	"xmin": -2,
	"xmax": 2,
	"ymin": -2,
	"ymax": 2,
	"zoom": 1,
}

func Man_handler(w http.ResponseWriter, r *http.Request) {
	for name := range params {
		s := r.FormValue(name)
		if s == "" {
			continue
		}
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			http.Error(w, fmt.Sprintf("query param %s: %s", name, err), http.StatusBadRequest)
			return
		}
		params[name] = f
	}
	if params["xmax"] <= params["xmin"] || params["ymax"] <= params["ymin"] {
		http.Error(w, "min coordinate greater than max", http.StatusBadRequest)
		return
	}
	xmin := params["xmin"]
	xmax := params["xmax"]
	ymin := params["ymin"]
	ymax := params["ymax"]
	zoom := params["zoom"]

	lenX := xmax - xmin
	midX := xmin + lenX/2
	xmin = midX - lenX/2/zoom
	xmax = midX + lenX/2/zoom
	lenY := ymax - ymin
	midY := ymin + lenY/2
	ymin = midY - lenY/2/zoom
	ymax = midY + lenY/2/zoom

	img := image.NewRGBA(image.Rect(0, 0, m_width, m_height))
	for py := 0; py < m_height; py++ {
		y := float64(py)/m_height*(ymax-ymin) + ymin
		for px := 0; px < m_width; px++ {
			x := float64(px)/m_width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	err := png.Encode(w, img)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	// const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return getColor(n)
		}
	}
	return color.Black
}

func getColor(n uint8) color.Color {
	paletted := [16]color.Color{
		color.RGBA{66, 30, 15, 255},    // # brown 3
		color.RGBA{25, 7, 26, 255},     // # dark violett
		color.RGBA{9, 1, 47, 255},      //# darkest blue
		color.RGBA{4, 4, 73, 255},      //# blue 5
		color.RGBA{0, 7, 100, 255},     //# blue 4
		color.RGBA{12, 44, 138, 255},   //# blue 3
		color.RGBA{24, 82, 177, 255},   //# blue 2
		color.RGBA{57, 125, 209, 255},  //# blue 1
		color.RGBA{134, 181, 229, 255}, // # blue 0
		color.RGBA{211, 236, 248, 255}, // # lightest blue
		color.RGBA{241, 233, 191, 255}, // # lightest yellow
		color.RGBA{248, 201, 95, 255},  // # light yellow
		color.RGBA{255, 170, 0, 255},   // # dirty yellow
		color.RGBA{204, 128, 0, 255},   // # brown 0
		color.RGBA{153, 87, 0, 255},    // # brown 1
		color.RGBA{106, 52, 3, 255},    // # brown 2
	}
	return paletted[n%16]
}
