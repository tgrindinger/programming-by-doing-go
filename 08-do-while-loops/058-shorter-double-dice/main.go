package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func RollDice(random *rand.Rand, writer io.Writer) {
	fmt.Fprintln(writer, "HERE COMES THE DICE!")
	fmt.Fprintln(writer)
	var roll1, roll2 int
	for ok := true; ok; ok = roll1 != roll2 {
		fmt.Fprintln(writer)
		roll1 = 1 + random.Intn(6)
		roll2 = 1 + random.Intn(6)
		fmt.Fprintf(writer, "Roll #1: %d\n", roll1)
		fmt.Fprintf(writer, "Roll #2: %d\n", roll2)
		fmt.Fprintf(writer, "The total is %d!\n", roll1 + roll2)
	}
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	RollDice(random, os.Stdout)
}
