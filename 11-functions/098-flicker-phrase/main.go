package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func first(writer io.Writer) {
	fmt.Fprint(writer, "I                               \r")
}

func second(writer io.Writer) {
	fmt.Fprint(writer, "  pledge                        \r")
}

func third(writer io.Writer) {
	fmt.Fprint(writer, "         allegiance             \r")
}

func fourth(writer io.Writer) {
	fmt.Fprint(writer, "                    to the      \r")
}

func fifth(writer io.Writer) {
	fmt.Fprint(writer, "                           flag.\r")
}

func PrintPhrase(random *rand.Rand, writer io.Writer) {
	for i := 0; i < 100000; i++ {
		r := 1 + random.Intn(5)
		if r == 1 {
			first(writer)
		} else if r == 2 {
			second(writer)
		} else if r == 3 {
			third(writer)
		} else if r == 4 {
			fourth(writer)
		} else if r == 5 {
			fifth(writer)
		}
	}
	fmt.Fprintln(writer, "I pledge allegiance to the flag.")
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	PrintPhrase(random, os.Stdout)
}
