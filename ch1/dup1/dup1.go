package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	input := bufio.NewScanner(os.Stdin)
	counts := make(map[string]int)

	for input.Scan() {
		counts[input.Text()]++
	}

	for key, value := range counts {
		fmt.Printf("%d\t%s\n", value, key)
	}
}