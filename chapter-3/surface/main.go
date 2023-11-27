// Вычисляет SVG-представление трехмерного графика функции
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 600, 320            // Размеры канвы в пикселях
	cells         = 200                 // Количество ячеек сетки
	xyrange       = 30.0                // Диапазон осей (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // Пикселей в единице x или y
	zscale        = height * 0.4        // Пикселей в единице z
	angle         = math.Pi / 6         // Углы осей x, y (=30°)
	accMidLimit   = 0.9                 // Значение коэффициента точности для раскраски желтым
	accHiLimit    = 0.95                // Значение коэффициента точности для раскраски синим и красным
)

func main() {
	http.HandleFunc("/", svgSurface) // Каждый запрос вызывает обработчик
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func svgSurface(w http.ResponseWriter, r *http.Request) { // http://localhost:8080/?mode=1&width=600&height=320&color=white
	var (
		mode           = 1 // 1 для холма с волнами, 2 для холма без волн, 3 для коробки для яиц
		width  float64 = 600
		height float64 = 320
		color          = "white"
	)

	modeUrl, err := strconv.Atoi(r.URL.Query().Get("mode"))
	if err != nil || modeUrl <= 0 {
		log.Println("mode from http ignored")
	} else {
		mode = modeUrl
	}

	widthUrl, err := strconv.ParseFloat(r.URL.Query().Get("width"), 64)
	if err != nil || widthUrl <= 0 {
		log.Println("width from http ignored")
	} else {
		width = widthUrl
	}

	heightUrl, err := strconv.ParseFloat(r.URL.Query().Get("height"), 64)
	if err != nil || heightUrl <= 0 {
		log.Println("height from http ignored")
	} else {
		height = heightUrl
	}

	colorUrl := r.URL.Query().Get("color")
	if err != nil || colorUrl == "" {
		log.Println("color from http ignored")
	} else {
		color = colorUrl
	}

	w.Header().Set("ContentType", "image/svg+xml")

	// Формируем HTML-страницу с SVG-изображением
	fmt.Fprintf(w, "<!DOCTYPE html>")
	fmt.Fprintf(w, "<html>")
	fmt.Fprintf(w, "<body>")

	// Вставляем SVG изображение
	surface(w, mode, width, height, color)

	fmt.Fprintf(w, "</body>")
	fmt.Fprintf(w, "</html>")
}

func surface(out io.Writer, mode int, width, height float64, color string) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' width='%d' height='%d'>",
		int(width), int(height))

	minZ, maxZ := findMinMaxZ(mode)
	if minZ == 0. { //nolint:gomnd
		minZ = maxZ * (1 - accHiLimit)
	}

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j, mode, width, height)
			bx, by, bz := corner(i, j, mode, width, height)
			cx, cy, cz := corner(i, j+1, mode, width, height)
			dx, dy, dz := corner(i+1, j+1, mode, width, height)

			if az > maxZ*accHiLimit || bz > maxZ*accHiLimit || cz > maxZ*accHiLimit || dz > maxZ*accHiLimit {
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='red'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
			} else if az > maxZ*accMidLimit || bz > maxZ*accMidLimit || cz > maxZ*accMidLimit || dz > maxZ*accMidLimit {
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='yellow'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
			} else if az < minZ*accHiLimit || bz < minZ*accHiLimit || cz < minZ*accHiLimit || dz < minZ*accHiLimit {
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='blue'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
			} else {
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(i, j, mode int, width, height float64) (sx, sy, z float64) {
	var (
		sin30 = math.Sin(angle) // sin(30°)
		cos30 = math.Cos(angle) // cos(30°)
	)

	// Ищем угловую точку (x,y) у ячейки (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Делаем проверку, при которой результат будет от -Inf до +Inf
	if x == 1 || y == 1 {
		return 0, 0, 0
	}

	// Вычисляем высоту поверхности z
	z = f(x, y, mode)

	// Изометрически проецируем (x,y,z) на двумерную канву SVG (sx,sy)
	sx = width/2 + (x-y)*cos30*xyscale

	if mode == 3 {
		sy = height/1.2 + (x+y)*sin30*xyscale - z*zscale
	} else {
		sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	}
	return sx, sy, z
}

func f(x, y float64, mode int) float64 {
	var r float64

	switch mode {
	case 1:
		r = math.Hypot(x, y) // Расстояние от (0,0)
		r = math.Sin(r) / r
	case 2:
		r = math.Exp(-(x*x + y*y)) // Холм
	case 3:
		r = math.Hypot(math.Cos(x), math.Sin(y)) // Коробка для яиц
		r = math.Sin(r) / r
	}
	return r
}

func cornerZ(i, j, mode int) (z float64) {
	// Ищем угловую точку (x,y) у ячейки (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Вычисляем высоту поверхности z
	z = f(x, y, mode)

	return z
}

func findMinMaxZ(mode int) (minZ, maxZ float64) {
	var allZ []float64

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			allZ = append(allZ, cornerZ(i+1, j, mode), cornerZ(i, j, mode), cornerZ(i, j+1, mode), cornerZ(i+1, j+1, mode))
		}
	}
	for _, z := range allZ {
		if z < minZ {
			minZ = z
		}

		if z > maxZ {
			maxZ = z
		}
	}
	return minZ, maxZ
}
