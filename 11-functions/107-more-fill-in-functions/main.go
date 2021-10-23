package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func superrand(random *rand.Rand, a, b int) int {
	var c int
	if a < b {
		c = a + random.Intn(b - a + 1)
	} else if a > b {
		c = b + random.Intn(a - b + 1)
	} else {
		c = a
	}
	return c
}

func stepcount(writer io.Writer, first, last, step int) {
	if first < last {
		for x := first; x <= last; x += step {
			fmt.Fprintf(writer, "%d ", x)
		}
	} else {
		for x := first; x >= last; x -= step {
			fmt.Fprintf(writer, "%d ", x)
		}
	}
}

func projectGrade(p1, p2, p3, p4, p5, p6 int) int {
	return (p1 * 6) +
		(p2 * 6) +
		(p3 * 4) +
		(p4 * 2) +
		(p5 * 1) +
		(p6 * 1)
}

func getName(reader io.Reader, writer io.Writer) string {
	scanner := bufio.NewScanner(reader)
	fmt.Fprintf(writer, "Please enter your name: ")
	scanner.Scan()
	return scanner.Text()
}

func crash(random *rand.Rand, writer io.Writer) {
	var magicWord string
	if random.Intn(10) == 0 {
		magicWord = "win"
	} else {
		magicWord = "lose"
	}
	fmt.Fprintf(writer, "you %s", magicWord)
}

func PrintFunctions(random *rand.Rand, reader io.Reader, writer io.Writer) {
	fmt.Fprintln(writer, "Here we go.")

	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Some random numbers, if you please: ")
	lo := 1
	hi := 10
	val1 := superrand(random, lo, hi)
	fmt.Fprintf(writer, "First: %d\n", val1)
	val2 := superrand(random, hi, lo)  // this time, put hi first
	fmt.Fprintf(writer, "Second: %d\n", val2)
	if val1 == val2 {
		fmt.Fprintln(writer, "Hey!  They were the same!")
	} else {
		fmt.Fprintln(writer, "They were not the same.")
	}

	fmt.Fprintln(writer)
	fmt.Fprint(writer, "More counting, but this time not by ones: ")
	// count from 2 to 8 by 2s
	stepcount(writer, 2, 8, 2)
	// count from 10 to 2 by 2s
	stepcount(writer, 10, 2, 2)

	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Let's figure a project grade.")
	a := 4
	b := 3
	c := 4
	d := 5
	e := 2
	f := 1
	result := projectGrade(a, b, c, d, e, f)
	fmt.Fprintf(writer, "434521 -> %d\n", result)

	fmt.Fprintln(writer)
	fmt.Fprint(writer, "Finally, some easy ones.")

	nombre := getName(reader, writer)
	fmt.Fprintf(writer, "Hi, %s\n", nombre)

	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Do you feel lucky, punk?")
	crash(random, writer)
	fmt.Fprintln(writer)
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	PrintFunctions(random, os.Stdin, os.Stdout)
}
