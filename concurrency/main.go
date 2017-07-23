package main

import "fmt"

func printNumbers() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}

func printLetters() {
	for i := 'A'; i < 'A'+10; i++ {
		fmt.Printf("%c ", i)
	}
}

func printOne() {
	printNumbers()
	printLetters()
}

func main() {
}
