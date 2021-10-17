package main

import (
	"fmt"
	"io"
	"os"
)

func PrintIfs(writer io.Writer) {
	people := 20
	cats := 30
	dogs := 15
	if people < cats {
		fmt.Fprintln(writer, "Too many cats!  The world is doomed!")
	}
	if people > cats {
		fmt.Fprintln(writer, "Not many cats!  The world is saved!")
	}
	if people < dogs {
		fmt.Fprintln(writer, "The world is drooled on!")
	}
	if people > dogs {
		fmt.Fprintln(writer, "The world is dry!")
	}
	dogs += 5
	if people >= dogs {
		fmt.Fprintln(writer, "People are greater than or equal to dogs.")
	}
	if people <= dogs {
		fmt.Fprintln(writer, "People are less than or equal to dogs.")
	}
	if people == dogs {
		fmt.Fprintln(writer, "People are dogs.")
	}
}

func main() {
	PrintIfs(os.Stdout)
}