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
	enemy := "STARMIE"
	fmt.Fprintf(writer, "Misty sent out %s!\n", enemy)
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "%s\t\tLv21\n", enemy)
	fmt.Fprintln(writer, "   HP [================]")
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "What will %s do?\n", pokeParty[0])
	fmt.Fprintln(writer, "\t1. Fight")
	fmt.Fprintln(writer, "\t2. Swap Pokemon")
	fmt.Fprintln(writer, "\t3. Run")
	fmt.Fprint(writer, "> ")
	scanner.Scan()
	choice, _ := strconv.Atoi(scanner.Text())
	if choice == 1 {
		fmt.Fprintln(writer)
		fmt.Fprintf(writer, "%s used THUNDERSHOCK!  It's super effective.\n", pokeParty[0])
		fmt.Fprintln(writer)
		fmt.Fprintf(writer, "%s\t\t Lv21\n", enemy)
		fmt.Fprintln(writer, "   HP [===========     ]")
		fmt.Fprintln(writer)
		fmt.Fprintf(writer, "%s used WATER PULSE!\n", enemy)
		fmt.Fprintf(writer, "%s fainted.\n", pokeParty[0])
	} else if choice == 2 {
		fmt.Fprintln(writer)
		fmt.Fprintf(writer, "%s swaps out with %s!\n", pokeParty[0], pokeParty[3])
		temp := pokeParty[0]
		pokeParty[0] = pokeParty[3]
		pokeParty[3] = temp
		fmt.Fprintln(writer)
		fmt.Fprintf(writer, "%s used BITE!  It's super effective.\n", pokeParty[0])
		fmt.Fprintln(writer)
		fmt.Fprintf(writer, "%s\t\t Lv21\n", enemy)
		fmt.Fprintln(writer, "   HP [                ]")
		fmt.Fprintln(writer)
		fmt.Fprintf(writer, "%s fainted.\n", enemy)
		fmt.Fprintf(writer, "%s gained 155 EXP. Points!\n", pokeParty[0])
		fmt.Fprintf(writer, "%s gained 155 EXP. Points!\n", pokeParty[3])
		fmt.Fprintln(writer, "Received TM03 - WATER PULSE.")
		fmt.Fprintln(writer, "Received CASCADEBADGE.")
		fmt.Fprintln(writer, "Received $2100.")
	} else if choice == 3 {
		fmt.Fprintln(writer)
		fmt.Fprintf(writer, "%s couldn't run. Skip a turn, coward.\n", pokeParty[0])
		fmt.Fprintln(writer)
		fmt.Fprintf(writer, "%s\t\t Lv21\n", enemy)
		fmt.Fprintln(writer, "   HP [================]")
		fmt.Fprintln(writer)
		fmt.Fprintf(writer, "%s used WATER PULSE!\n", enemy)
		fmt.Fprintf(writer, "%s fainted.\n", pokeParty[0])
	}
}

func main() {
	PlayGame(os.Stdin, os.Stdout)
}
