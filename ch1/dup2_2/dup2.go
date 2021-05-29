// dup2 using map[string]map[string]int to print the name of the file in the output
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	files := os.Args[1:]
	counts := map[string]map[string]int{}

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for file, lines := range counts {
		for line, n := range lines {
			if n > 1 {
				fmt.Printf("%s: %d\t%s\n", file, n, line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int){
	input := bufio.NewScanner(f)
	counts[f.Name()] = make(map[string]int)
	for input.Scan() {
		counts[f.Name()][input.Text()]++
	}
}