package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func GuessNumber(random *rand.Rand, reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	answer := 1 + random.Intn(10)
	fmt.Fprintln(writer, "I'm thinking of a number from 1 to 10.")
	fmt.Fprint(writer, "Your guess: ")
	scanner.Scan()
	guess, _ := strconv.Atoi(scanner.Text())
	fmt.Fprintln(writer)
	if guess == answer {
		fmt.Fprintf(writer, "That's right!  My secret number was %d!\n", answer)
	} else {
		fmt.Fprintf(writer, "Sorry, but I was really thinking of %d.\n", answer)
	}
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	GuessNumber(random, os.Stdin, os.Stdout)
}
