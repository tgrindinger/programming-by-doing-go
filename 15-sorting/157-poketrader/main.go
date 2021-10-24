package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func PlayGame(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	pokeParty := []string{ "PIKACHU", "CHARMELEON", "GEODUDE", "GYARADOS", "BUTTERFREE", "MANKEY" }
	var x int
	for ok := true; ok; ok = x > 0 {
		fmt.Fprintln(writer, "EXCHANGE POKEMON")
		fmt.Fprintln(writer)
		fmt.Fprintf(writer, "0. %s\n", pokeParty[0])
		for i := 1; i < len(pokeParty); i++ {
			fmt.Fprintf(writer, "\t%d. %s\n", i, pokeParty[i])
		}
		fmt.Fprintln(writer)
		fmt.Fprintf(writer, "Choose a Pokemon to exchange with %s. (Or enter 0 to quit.)\n", pokeParty[0])
		fmt.Fprint(writer, "> ")
		scanner.Scan()
		x, _ = strconv.Atoi(scanner.Text())

		if x != 0 {
			temp := pokeParty[0]
			pokeParty[0] = pokeParty[x]
			pokeParty[x] = temp
		}
	}
}

func main() {
	PlayGame(os.Stdin, os.Stdout)
}
