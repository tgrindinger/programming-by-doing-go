package main

import (
	"fmt"
	"io"
	"os"
)

func PrintVariables(writer io.Writer) {
	var name, eyes, teeth, hair string
	var age, height, weight int
	name = "Zed A. Shaw"
	age = 35
	height = 74
	weight = 180
	eyes = "Blue"
	teeth = "White"
	hair = "Brown"
	fmt.Fprintf(writer, "Let's talk about %s.\n", name)
	fmt.Fprintf(writer, "He's %d inches tall.\n", height)
	fmt.Fprintf(writer, "He's %d pounds heavy.\n", weight)
	fmt.Fprintf(writer, "Actually that's not too heavy.\n")
	fmt.Fprintf(writer, "He's got %s eyes and %s hair.\n", eyes, hair)
	fmt.Fprintf(writer, "His teeth are usually %s depending on the coffee.\n", teeth)
	fmt.Fprintf(writer, "If I add %d, %d, and %d I get %d.\n", age, height, weight, age + height + weight)
}

func main() {
	PrintVariables(os.Stdout)
}