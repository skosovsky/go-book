// Генерирует анимированный gif из случайных фигур Лиссажу
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

func main() {
	lissaj(os.Stdout)
}

func lissaj(out io.Writer) {
	var pallete = []color.Color{color.Black, color.White,
		color.RGBA{R: 60, G: 179, B: 113, A: 1},
		color.RGBA{R: 255, G: 0, B: 0, A: 1},
		color.RGBA{R: 255, G: 255, B: 0, A: 1},
		color.RGBA{R: 0, G: 255, B: 255, A: 1},
		color.RGBA{R: 255, G: 0, B: 255, A: 1}}

	const (
		cycles  = 5     // Количество полных колебаний x
		res     = 0.001 // Угловое разрешение
		size    = 100   // Канва изображения охватывает [-size..+size]
		nframes = 64    // Количество кадров анимации
		delay   = 8     // Задержка между кадрами (единица - 10 мс)
	)
	rand.New(rand.NewSource(time.Now().UnixNano()))
	freq := rand.Float64() * 3.0 // Относительная частота колебаний y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Разность фаз

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		colorIndex := uint8(rand.Intn(len(pallete))) // генерируем случайный цвет
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
