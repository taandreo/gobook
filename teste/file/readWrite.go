package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

const readFilename = "read.txt"
const writeFilename = "write.txt"

func main(){
	writeFile()
}

func readAllFile(){
	// ioutil.ReadFile
	// le o arquivo inteiro na memória, eficiente para pequenos arquivos.
	data, err := ioutil.ReadFile(readFilename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
	}
	fmt.Println(string(data))
}

func readLines(){
	// Reading a file line by line
	file, err := os.Open(readFilename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func writeFile(){
	file, err := os.Create(writeFilename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer file.Close()
	s := "Alecrim, alecrim dourado\nQue nasceu no campo sem ser semeado.\n"
	file.WriteString(s)
}