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
	answer := 1 + random.Intn(100)
	fmt.Fprintln(writer, "I'm thinking of a number between 1-100.  Try to guess it.")
	fmt.Fprint(writer, "> ")
	scanner.Scan()
	guess, _ := strconv.Atoi(scanner.Text())
	if guess < answer {
		fmt.Fprintf(writer, "Sorry, you are too low.  I was thinking of %d.\n", answer)
	} else if guess > answer {
		fmt.Fprintf(writer, "Sorry, you are too high.  I was thinking of %d.\n", answer)
	} else {
		fmt.Fprintf(writer, "You guessed it!  What are the odds?!?")
	}
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	GuessNumber(random, os.Stdin, os.Stdout)
}
