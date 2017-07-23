package main

import (
	"testing"
	"time"
)

func TestPrintOne(t *testing.T) {
	printOne()
}

func TestGoPrintOne(t *testing.T) {
	goPrintOne()
	time.Sleep(1 * time.Millisecond)
}

func BenchmarkPrintOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printOne()
	}
}

func BenchmarkGoPrintOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrintOne()
	}
}
