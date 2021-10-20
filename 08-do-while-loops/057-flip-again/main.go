package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func Flip(random *rand.Rand, reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	var again string
	for ok := true; ok; ok = again == "y" {
		flip := random.Intn(2)
		var coin string
		if flip == 1 {
			coin = "HEADS"
		} else {
			coin = "TAILS"
		}
		fmt.Fprintf(writer, "You flip a coin and it is... %s\n", coin)
		fmt.Fprint(writer, "Would you like to flip again (y/n)? ")
		scanner.Scan()
		again = scanner.Text()
	}
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	Flip(random, os.Stdin, os.Stdout)
}
