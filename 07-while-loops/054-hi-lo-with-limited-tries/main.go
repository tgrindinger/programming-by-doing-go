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

func PlayHiLo(random *rand.Rand, reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	answer := 1 + random.Intn(100)
	tries := 1
	fmt.Fprintln(writer, "I'm thinking of a number between 1-100.  You have 7 guesses.")
	fmt.Fprint(writer, "First guess: ")
	scanner.Scan()
	guess, _ := strconv.Atoi(scanner.Text())
	for guess != answer && tries < 7 {
		tries++
		if guess < answer {
			fmt.Fprintln(writer, "Sorry, you are too low.")
		} else {
			fmt.Fprintln(writer, "Sorry, you are too high.")
		}
		fmt.Fprintf(writer, "Guess # %d: ", tries)
		scanner.Scan()
		guess, _ = strconv.Atoi(scanner.Text())
	}
	if guess == answer {
		fmt.Fprintln(writer, "You guessed it!  What are the odds?!?")
	} else {
		fmt.Fprintln(writer, "Sorry, you didn't guess it in 7 tries.  You lose.")
	}
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	PlayHiLo(random, os.Stdin, os.Stdout)
}
