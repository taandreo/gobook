package main

import (
	"fmt"
	"os"
	"strings"
	"io/ioutil"
)

func main(){
	files := os.Args[1:]
	counts := make(map[string]int)

	for _, fname := range files {
		data, err := ioutil.ReadFile(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

