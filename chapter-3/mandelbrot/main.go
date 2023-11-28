// // Создает PNG-изображение фрактала Мандельброта
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", pngMandelbrot) // Каждый запрос вызывает обработчик
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func pngMandelbrot(w http.ResponseWriter, r *http.Request) { // http://localhost:8080/?mode=1&minCoef=2&maxCoef=2&width=1024&height=1024

	var (
		xMin, yMin, xMax, yMax = -2, -2, +2, +2
		width, height          = 1024, 1024
		minCoef, maxCoef       = 1, 1
		mode                   = 1 // от 1 до 5, разные режимы
	)

	modeUrl, err := strconv.Atoi(r.URL.Query().Get("mode"))
	if err != nil || modeUrl <= 0 {
		log.Println("mode from http ignored")
	} else {
		mode = modeUrl
	}

	minCoefUrl, err := strconv.Atoi(r.URL.Query().Get("minCoef"))
	if err != nil || minCoefUrl <= 0 {
		log.Println("minCoef from http ignored")
	} else {
		minCoef = minCoefUrl
	}

	maxCoefUrl, err := strconv.Atoi(r.URL.Query().Get("maxCoef"))
	if err != nil || maxCoefUrl <= 0 {
		log.Println("maxCoef from http ignored")
	} else {
		maxCoef = maxCoefUrl
	}

	widthUrl, err := strconv.Atoi(r.URL.Query().Get("width"))
	if err != nil || widthUrl <= 0 {
		log.Println("width from http ignored")
	} else {
		width = widthUrl
	}

	heightUrl, err := strconv.Atoi(r.URL.Query().Get("height"))
	if err != nil || heightUrl <= 0 {
		log.Println("height from http ignored")
	} else {
		height = heightUrl
	}

	xMin, yMin, xMax, yMax = xMin*minCoef, yMin*minCoef, xMax*maxCoef, yMax*maxCoef
	// Создаем PNG изображение
	img := renderPng(mode, xMin, yMin, xMax, yMax, width, height)

	w.Header().Set("Content-Type", "image/png")

	if err := png.Encode(w, img); err != nil {
		log.Println(err)
	}
}

func renderPng(mode, xMin, yMin, xMax, yMax, width, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for pY := 0; pY < height; pY++ {
		y := float64(pY)/float64(height)*(float64(yMax)-float64(yMin)) + float64(yMin)
		for pX := 0; pX < width; pX++ {
			x := float64(pX)/float64(width)*(float64(xMax)-float64(xMin)) + float64(xMin)
			z := complex(x, y)
			// Точка (pX, pY) представляет комплексное значение z
			switch mode {
			case 1:
				img.Set(pX, pY, mandelbrot(z))
			case 2:
				img.Set(pX, pY, mandelbrotColor(z))
			case 3:
				img.Set(pX, pY, acos(z))
			case 4:
				img.Set(pX, pY, sqrt(z))
			case 5:
				img.Set(pX, pY, newton(z))
			default:
				img.Set(pX, pY, mandelbrot(z))
			}
		}
	}
	return img
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{Y: 255 - contrast*n}
		}
	}
	return color.Black
}

func mandelbrotColor(z complex128) color.Color {
	const iterations = 200

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			blue := uint8(real(v)*128) + 127
			red := uint8(imag(v)*128) + 127
			return color.YCbCr{Y: 192, Cb: blue, Cr: red}
		}
	}
	return color.Black
}

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{Y: 192, Cb: blue, Cr: red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{Y: 128, Cb: blue, Cr: red}
}

func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 { //nolint:gomnd
			return color.Gray{Y: 255 - contrast*i}
		}
	}
	return color.Black
}
