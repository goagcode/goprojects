package main

import (
	"time"
)

func printNumbers() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
	}
}

func printLetters() {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
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
