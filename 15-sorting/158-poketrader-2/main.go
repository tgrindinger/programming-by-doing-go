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
	var a int
	for ok := true; ok; ok = a >= 0 {
		fmt.Fprintln(writer, "EXCHANGE POKEMON")
		fmt.Fprintln(writer)
		for i := 0; i < len(pokeParty); i++ {
			fmt.Fprintf(writer, "\t%d. %s\n", i, pokeParty[i])
		}
		fmt.Fprintln(writer)
		fmt.Fprint(writer, "Choose a Pokemon (Or -1 to quit). => ")
		scanner.Scan()
		a, _ = strconv.Atoi(scanner.Text())
		if a >= 0 {
			fmt.Fprintf(writer, "Choose a Pokemon to exchange with %s => ", pokeParty[a])
			scanner.Scan()
			b, _ := strconv.Atoi(scanner.Text())
			fmt.Fprintf(writer, "Exchanging %s with %s.\n", pokeParty[a], pokeParty[b])
			temp := pokeParty[a]
			pokeParty[a] = pokeParty[b]
			pokeParty[b] = temp
		}
	}
}

func main() {
	PlayGame(os.Stdin, os.Stdout)
}
