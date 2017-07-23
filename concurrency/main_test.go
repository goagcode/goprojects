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
