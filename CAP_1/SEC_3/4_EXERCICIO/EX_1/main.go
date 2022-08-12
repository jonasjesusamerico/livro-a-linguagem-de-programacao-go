package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, "os.Stdin", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Println(fmt.Fprintf(os.Stderr, "dup2: %v", err))
				continue
			}
			countLines(f, arg, counts)
			f.Close()
		}
	}

	for line, filenames := range counts {
		fileCount := len(filenames)
		if fileCount == 1 {
			total := 0
			for _, count := range filenames {
				total += count
			}
			if total <= 1 {
				continue
			}
		}
		fmt.Printf("Encontrado em %d arquivo \t%s\n", fileCount, line)

		for name, count := range filenames {
			fmt.Printf("\t%d vezes em %s\n", count, name)
		}
	}
}

func countLines(f *os.File, filename string, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		fmt.Println(input.Text())
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][filename]++
	}
}
