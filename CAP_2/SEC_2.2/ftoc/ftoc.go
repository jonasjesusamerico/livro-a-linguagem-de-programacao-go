package main

import "fmt"

func main() {
	const freezigF, boilingF = 32.0, 212.0
	fmt.Printf("%g F = %g C\n", freezigF, fToc(freezigF))
	fmt.Printf("%g F = %g C\n", boilingF, fToc(boilingF))

}

func fToc(f float64) float64 {
	return (f - 32) * 5 / 9
}
