package main

import (
	"os"

	"github.com/zserge/webview"
)

func main() {
	// Open wikipedia in a 800x600 resizable window
	webview.Open("Minimal webview example", os.Args[1], 800, 600, true)
}
