package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func credits(writer io.Writer) {
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "programmed by Graham Mitchell")
	fmt.Fprintln(writer, "modified by Thomas Grindinger")
	fmt.Fprintln(writer, "This code is distributed under the terms of the standard BSD license.  Do with it as you wish.")
}

func randchar(random *rand.Rand) int {
	return 'A' + random.Intn(26)
}

func counter(writer io.Writer, start, stop int) {
	for ctr := start; ctr <= stop; ctr++ {
		fmt.Fprintf(writer, "%d ", ctr)
	}
}

func abso(value int) int {
	absval := value
	if value < 0 {
		absval = -value
	}
	return absval
}

func PrintFunctions(random *rand.Rand, writer io.Writer) {
	fmt.Fprintln(writer, "Watch as we demonstrate functions.")

	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "I'm going to get a random character from A-Z")
	c := randchar(random)
	fmt.Fprintf(writer, "The character is: %c\n", c)

	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Now let's count from -10 to 10")
	begin := -10
	end := 10
	counter(writer, begin, end)
	fmt.Fprintln(writer, "How was that?")

	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Now we take the absolute value of a number.")
	x := -10
	y := abso(x)
	fmt.Fprintf(writer, "|%d| = %d\n", x, y)

	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "That's all.  This program has been brought to you by:")
	credits(writer)
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	PrintFunctions(random, os.Stdout)
}
