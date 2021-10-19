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

func PlayGame(random *rand.Rand, reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	answer := 1 + random.Intn(3)
	fmt.Fprintln(writer, "You slide up to Fast Eddie's card table and plop down your cash.")
	fmt.Fprintln(writer, "He glances at you out of the corner of his eye and starts shuffling.")
	fmt.Fprintln(writer, "He lays down three cards.")
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Which one is the ace?")
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "        ##  ##  ##")
	fmt.Fprintln(writer, "        ##  ##  ##")
	fmt.Fprintln(writer, "        1   2   3")
	fmt.Fprintln(writer)
	fmt.Fprint(writer, "> ")
	scanner.Scan()
	guess, _ := strconv.Atoi(scanner.Text())
	fmt.Fprintln(writer)
	if guess == answer {
		fmt.Fprintln(writer, "You nailed it! Fast Eddie reluctantly hands over your winnings, scowling.")
	} else {
		fmt.Fprintf(writer, "Ha! Fast Eddie wins again! The ace was card number %d.\n", answer)
	}
	fmt.Fprintln(writer)
	if answer == 1 {
		fmt.Fprintln(writer, "        AA  ##  ##")
		fmt.Fprintln(writer, "        AA  ##  ##")
	} else if answer == 2 {
		fmt.Fprintln(writer, "        ##  AA  ##")
		fmt.Fprintln(writer, "        ##  AA  ##")
	} else {
		fmt.Fprintln(writer, "        ##  ##  AA")
		fmt.Fprintln(writer, "        ##  ##  AA")
	}
	fmt.Fprintln(writer, "        1   2   3")
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	PlayGame(random, os.Stdin, os.Stdout)
}
