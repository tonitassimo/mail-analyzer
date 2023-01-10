package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

const CHUNK_SIZE = 10

func main() {
	filename := processArgs(os.Args)
	processFile(filename)
}

func processArgs(args []string) string {
	if len(args) < 2 {
		fmt.Println("No file provided. Provide a filename as the first command line argument.")
		os.Exit(1)
	}
	return args[1]
}

func processFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	b := []string{}
	scanner := bufio.NewScanner(file)
	var wg sync.WaitGroup

	for scanner.Scan() {
		b = append(b, scanner.Text())
		if len(b) == CHUNK_SIZE {
			wg.Add(1)
			go processChunk(b, &wg)
			b = nil
		}
	}

	wg.Add(1)
	go processChunk(b, &wg)

	wg.Wait()
}

func processChunk(chunk []string, wg *sync.WaitGroup) {
	for i, v := range chunk {
		processLine(strconv.Itoa(i) + ".: " + v)
	}

	defer wg.Done()
}

func processLine(line string) {
	if validateLine(line) {
		fmt.Println(line)
	} else {
		fmt.Println("err: no valid email address!")
	}
}

func validateLine(line string) bool {
	return strings.Contains(line, "@")
}
