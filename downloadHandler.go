package main

import (
	"fmt"
	"os"

	"github.com/raderh2o/godl/tutil"
)

func makeHandler(file *os.File, t tutil.Terminal) func(int, []byte) {

	downloaded := 0
	return func(n int, b []byte) {
		downloaded += n
		t.ClearScreen()
		fmt.Printf("Downloaded %d bytes\n", downloaded)
		t.MoveCursorToHome()
		file.Write(b)
	}
}
