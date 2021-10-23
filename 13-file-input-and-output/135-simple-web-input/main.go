package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func PrintWebLine(writer io.Writer) {
	url := "https://programmingbydoing.com/a/examples/SimpleWebInput.java"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Unable to open url %q: %v", url, err)
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	scanner.Scan()
	fmt.Fprintln(writer, scanner.Text())
	scanner.Scan()
	fmt.Fprintln(writer, scanner.Text())
}

func main() {
	PrintWebLine(os.Stdout)
}
