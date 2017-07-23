package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
}

func printLetters() {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
}

func printOne() {
	printNumbers()
	printLetters()
}

func goPrintOne() {
	go printNumbers()
	go printLetters()
}

func main() {
}
