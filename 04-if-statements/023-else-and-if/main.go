package main

import (
	"fmt"
	"io"
	"os"
)

func PrintIfs(writer io.Writer) {
	people := 30
	cars := 40
	buses := 15
	if cars > people {
		fmt.Fprintln(writer, "We should take the cars.")
	} else if cars < people {
		fmt.Fprintln(writer, "We should not take the cars.")
	} else {
		fmt.Fprintln(writer, "We can't decide.")
	}
	if buses > cars {
		fmt.Fprintln(writer, "That's too many buses.")
	} else if buses < cars {
		fmt.Fprintln(writer, "Maybe we could take the buses.")
	} else {
		fmt.Fprintln(writer, "We still can't decide.")
	}
	if people > buses {
		fmt.Fprintln(writer, "All right, let's just take the buses.")
	} else {
		fmt.Fprintln(writer, "Fine, let's stay home then.")
	}
}

func main() {
	PrintIfs(os.Stdout)
}
