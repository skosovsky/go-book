// Минимальный "echo"-сервер
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
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var m sync.Mutex
var count int32

func main() {
	http.HandleFunc("/", handler) // Каждый запрос вызывает обработчик
	http.HandleFunc("/count", counter)
	http.HandleFunc("/gif", gifLissaj)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// Обработчик вызывает компонент по пути из url запроса
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s %s/n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

	atomic.AddInt32(&count, 1)
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	m.Unlock()
}

func gifLissaj(w http.ResponseWriter, r *http.Request) { // http://localhost:8080/gif?cycles=50&res=0.005&size=200&nframes=128&delay=16
	var (
		cycles  = 5     // Количество полных колебаний x
		res     = 0.001 // Угловое разрешение
		size    = 100   // Канва изображения охватывает [-size..+size]
		nframes = 64    // Количество кадров анимации
		delay   = 8     // Задержка между кадрами (единица - 10 мс)
	)

	cyclesUrl, err := strconv.Atoi(r.URL.Query().Get("cycles"))
	if err != nil || cyclesUrl <= 0 {
		log.Println("cycles from http ignored")
	} else {
		cycles = cyclesUrl
	}

	resUrl, err := strconv.ParseFloat(r.URL.Query().Get("res"), 64)
	if err != nil || resUrl <= 0 {
		log.Println("res from http ignored")
	} else {
		res = resUrl
	}

	sizeUrl, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil || sizeUrl <= 0 {
		log.Println("size from http ignored")
	} else {
		size = sizeUrl
	}

	nframesUrl, err := strconv.Atoi(r.URL.Query().Get("nframes"))
	if err != nil || nframesUrl <= 0 {
		log.Println("nframes from http ignored")
	} else {
		nframes = nframesUrl
	}

	delayUrl, err := strconv.Atoi(r.URL.Query().Get("delay"))
	if err != nil || delayUrl <= 0 {
		log.Println("delay from http ignored")
	} else {
		delay = delayUrl
	}

	lissaj(w, cycles, res, size, nframes, delay)
}

func lissaj(out io.Writer, cycles int, res float64, size int, nframes int, delay int) {

	var pallete = []color.Color{color.Black, color.White,
		color.RGBA{R: 60, G: 179, B: 113, A: 1},
		color.RGBA{R: 255, G: 0, B: 0, A: 1},
		color.RGBA{R: 255, G: 255, B: 0, A: 1},
		color.RGBA{R: 0, G: 255, B: 255, A: 1},
		color.RGBA{R: 255, G: 0, B: 255, A: 1}}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	freq := rand.Float64() * 3.0 // Относительная частота колебаний y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Разность фаз

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		colorIndex := uint8(rand.Intn(len(pallete))) // генерируем случайный цвет
		for t := 0.0; t < float64(cycles*2)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), colorIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
