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
	answer := 1 + random.Intn(10)
	scanner := bufio.NewScanner(reader)
	tries := 1
	fmt.Fprintln(writer, "I have chosen a number between 1 and 10. Try to guess it.")
	fmt.Fprint(writer, "Your guess: ")
	scanner.Scan()
	guess, _ := strconv.Atoi(scanner.Text())
	for guess != answer {
		fmt.Fprintln(writer, "That is incorrect.  Guess again.")
		fmt.Fprint(writer, "Your guess: ")
		scanner.Scan()
		guess, _ = strconv.Atoi(scanner.Text())
		tries++
	}
	fmt.Fprintln(writer, "That's right!  You're a good guesser.")
	fmt.Fprintf(writer, "It only took you %d tries.\n", tries)
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	GuessNumber(random, os.Stdin, os.Stdout)
}
