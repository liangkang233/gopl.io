package server_surface

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

const (
	// width, height = 600, 320            // canvas size in pixels
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

var params1 = map[string]float64{
	"index":  0,
	"height": 320,
	"width":  600,
}

func Svg_handler(w http.ResponseWriter, r *http.Request) {

	/* 	直接像lissajous那样用form参数判断不太好，用map查找
	   	if err := r.ParseForm(); err != nil {
	   		log.Print(err)
	   	}
	   	for k, v := range r.Form {
	   		var err error
	   		if k == "index" {
	   			index, err = strconv.Atoi(v[0]) //get发送的同名变量index只接受第一个
	   			if err != nil {
	   				fmt.Fprintf(os.Stderr, "service_surface str conv form: %v\n", err)
	   				os.Exit(1)
	   			}
	   		} else if k == "height" {
	   			height, err = strconv.Atoi(v[0]) //get发送的同名变量height只接受第一个
	   			if err != nil {
	   				fmt.Fprintf(os.Stderr, "service_surface str conv form: %v\n", err)
	   				os.Exit(1)
	   			}
	   		} else if k == "width" {
	   			width, err = strconv.Atoi(v[0]) //get发送的同名变量width只接受第一个
	   			if err != nil {
	   				fmt.Fprintf(os.Stderr, "service_surface str conv form: %v\n", err)
	   				os.Exit(1)
	   			}
	   		}
	   	} */
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
	if params["height"] < 0 || params["width"] < 0 {
		http.Error(w, "min coordinate greater than max", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", int(params1["width"]), int(params1["height"]))
	surface(w)
	fmt.Fprintln(w, "</svg>")
}

func surface(w http.ResponseWriter) {
	z_min, z_max := min_max()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err1 := corner(i+1, j)
			bx, by, err2 := corner(i, j)
			cx, cy, err3 := corner(i, j+1)
			dx, dy, err4 := corner(i+1, j+1)
			if err1 || err2 || err3 || err4 {
				fmt.Printf("区块%d %d，corner() 产生无效数值，忽略之\n", i, j)
				continue
			}
			fmt.Fprintf(w, "<polygon style='stroke: %s;' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				svgColor(i, j, z_min, z_max), ax, ay, bx, by, cx, cy, dx, dy) // 练习3.3 多边形上色
		}
	}
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	width, height := params1["width"], params1["height"]
	var xyscale = float64(width) / 2 / xyrange // pixels per x or y unit
	var zscale = float64(height) * 0.4         // pixels per z unit

	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	err := math.IsNaN(sx) || math.IsNaN(sy) //练习3.1
	return sx, sy, err
}

func f(x, y float64) float64 {
	switch params1["index"] {
	case 1: //鸡蛋盒形状
		r := 0.2 * (math.Cos(x) + math.Cos(y))
		return r
	case 2: //马鞍形状
		a := 25.0
		b := 17.0
		a2 := a * a
		b2 := b * b
		r := y*y/a2 - x*x/b2
		return r
	default: //水滴形状
		r := math.Hypot(x, y) // distance from (0,0)
		return math.Sin(r) / r
	}
}

// minmax返回给定x和y的最小值/最大值并假设为方域的z的最小值和最大值。
func min_max() (min, max float64) {
	min = math.NaN()
	max = math.NaN()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			for xoff := 0; xoff <= 1; xoff++ {
				for yoff := 0; yoff <= 1; yoff++ {
					x := xyrange * (float64(i+xoff)/cells - 0.5)
					y := xyrange * (float64(j+yoff)/cells - 0.5)
					z := f(x, y)
					if math.IsNaN(min) || z < min {
						min = z
					}
					if math.IsNaN(max) || z > max {
						max = z
					}
				}
			}
		}
	}
	return min, max
}

func svgColor(i, j int, zmin, zmax float64) string {
	min := math.NaN()
	max := math.NaN()
	for xoff := 0; xoff <= 1; xoff++ {
		for yoff := 0; yoff <= 1; yoff++ {
			x := xyrange * (float64(i+xoff)/cells - 0.5)
			y := xyrange * (float64(j+yoff)/cells - 0.5)
			z := f(x, y)
			if math.IsNaN(min) || z < min {
				min = z
			}
			if math.IsNaN(max) || z > max {
				max = z
			}
		}
	}

	// 红色（#ff0000）, 蓝色（#0000ff）
	color := ""
	if math.Abs(max) > math.Abs(min) { // 向下凸
		red := math.Exp(math.Abs(max)) / math.Exp(math.Abs(zmax)) * 255
		if red > 255 {
			red = 255
		}
		color = fmt.Sprintf("#%02x0000", int(red))
	} else { // 向上凹
		blue := math.Exp(math.Abs(min)) / math.Exp(math.Abs(zmin)) * 255
		if blue > 255 {
			blue = 255
		}
		color = fmt.Sprintf("#0000%02x", int(blue))
	}
	return color
}
