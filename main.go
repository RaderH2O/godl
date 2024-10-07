package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/raderh2o/godl/downloader"
	"github.com/raderh2o/godl/tutil"
)

func main() {
	t := tutil.Terminal{}
	t.ClearScreen()
	fmt.Print("godl - enter a link to download from >> ")
	// TODO: get an input (with `bufio`), send an HTTP request with the http package (using `http.Get()`), and save the file

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	url := scanner.Text() // NOTE: This is the user input for the url

	fmt.Print("godl - enter the filename you want to save the file as >> ")
	scanner.Scan()
	filename := scanner.Text() // NOTE: This is the user input for the filename

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	err = downloader.Download(url, 1_000_000, makeHandler(file, t))
	if err != nil {
		fmt.Println("Error downloading file:", err)
		return
	}
}
