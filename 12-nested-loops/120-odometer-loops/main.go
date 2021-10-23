package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func PrintCounter(reader io.Reader, writer io.Writer, ms int64) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprint(writer, "Which base (2-10): ")
	scanner.Scan()
	base, _ := strconv.Atoi(scanner.Text())
	for thous := 0; thous < base; thous++ {
		for hund := 0; hund < base; hund++ {
			for tens := 0; tens < base; tens++ {
				for ones := 0; ones < base; ones++ {
					fmt.Fprintf(writer, " %d%d%d%d\r", thous, hund, tens, ones)
					time.Sleep(time.Duration(ms * 1000))
				}
			}
		}
	}
	fmt.Fprintln(writer)
}

func main() {
	PrintCounter(os.Stdin, os.Stdout, 10)
}
