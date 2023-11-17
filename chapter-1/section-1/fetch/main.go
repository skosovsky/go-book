// Выводи ответ на запрос по заданному url
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		switch {
		case strings.HasPrefix(url, "http://"):
			return
		case strings.HasPrefix(url, "https://"):
			return
		default:
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Println("код состояния:", resp.Status)
		io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
		}
		fmt.Println("")
		
		//b, err := io.ReadAll(resp.Body)
		//resp.Body.Close()
		//if err != nil {
		//	fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
		//	os.Exit(1)
		//}
		//fmt.Printf("%s", b)
	}
}
