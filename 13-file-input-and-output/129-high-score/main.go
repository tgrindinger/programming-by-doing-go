package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func SaveHighScore(reader io.Reader, writer io.Writer, fileWriter io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprintln(writer, "You got a high score!")
	fmt.Fprintln(writer)
	fmt.Fprint(writer, "Please enter your score: ")
	scanner.Scan()
	score, _ := strconv.Atoi(scanner.Text())
	fmt.Fprint(writer, "Please enter your name: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Fprintf(fileWriter, "%s %d\n", name, score)
	fmt.Fprintln(writer, "Data stored into score.txt.")
}

func main() {
	writer, _ := os.Create("score.txt")
	SaveHighScore(os.Stdin, os.Stdout, writer)
	writer.Close()
}
