package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func BenchmarkSemJoin(b *testing.B) {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func BenchmarkJoin(b *testing.B) {
	fmt.Println(strings.Join(os.Args[1:], ""))
}
