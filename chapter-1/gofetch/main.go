// Выполняет параллельную выборку url и сообщает о затраченном времени и размере ответа для каждого из них
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	var m sync.Mutex

	f, err := os.Create("fetch.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	for _, url := range os.Args[1:] {
		switch {
		case strings.HasPrefix(url, "http://"):
		case strings.HasPrefix(url, "https://"):
		default:
			url = "http://" + url
		}

		go fetch(url, ch, m, *f) // Запуск горутины
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // Получение из канала ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, m sync.Mutex, f os.File) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // Отправка в канал ch
		return
	}
	defer resp.Body.Close() // Исключение утечки ресурсов

	m.Lock()
	nbytes, err := io.Copy(&f, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v, url, err")
		return
	}
	m.Unlock()

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
