// Fetchall busca URLs em paralelo e informa os tempos gastos e os tamanhos
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)


const filename = "output.txt"

func main(){
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // inicia um gorotine
	}
	file, err := os.Create(filename)
	check(err)
	defer file.Close()

	for range os.Args[1:] {
		out := <-ch // recebe do canal ch
		fmt.Println(out)
		file.WriteString(out + "\n")
	}
	out := fmt.Sprintf("%.2fs elapesed\n", time.Since(start).Seconds())
	fmt.Print(out)
	file.WriteString(out)
}

func fetch(url string, ch chan<- string){
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // envia para o canal ch
		return
	}
	
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // evita vzamento de recursos
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func check(err error){
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}