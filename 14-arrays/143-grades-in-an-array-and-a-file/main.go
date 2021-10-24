package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func PrintGrades(random *rand.Rand, reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprint(writer, "Name (first last): ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Fprint(writer, "Filename: ")
	scanner.Scan()
	filename := scanner.Text()
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Here are your randomly-selected grades (hope you pass):")
	arr := make([]int, 5)
	for i := 0; i < len(arr); i++ {
		arr[i] = 1 + random.Intn(100)
		fmt.Fprintf(writer, "Grade %d: %d\n", i + 1, arr[i])
	}
	fmt.Fprintln(writer)
	file, _ := os.Create(filename)
	defer file.Close()
	fmt.Fprintln(file, name)
	for i := 0; i < len(arr); i++ {
		fmt.Fprintf(file, "%d ", arr[i])
	}
	fmt.Fprintln(file)
	fmt.Fprintf(writer, "Data saved in %q.\n", filename)
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	PrintGrades(random, os.Stdin, os.Stdout)
}
