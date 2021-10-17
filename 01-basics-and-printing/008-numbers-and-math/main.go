package main

import (
	"fmt"
	"io"
	"os"
)

func PrintMath(out io.Writer) {
	fmt.Fprintln(out, "I will now count my chickens:")
	fmt.Fprintln(out, "Hens", 25+30/6)
	fmt.Fprintln(out, "Roosters", 100-25*3%4)
	fmt.Fprintln(out, "Now I will count the eggs:")
	fmt.Fprintln(out, 3+2+1-5+4%2-1/4+6)
	fmt.Fprintln(out, "Is it true that 3 + 2 < 5 - 7?")
	fmt.Fprintln(out, 3+2 < 5-7)
	fmt.Fprintln(out, "What is 3 + 2?", 3+2)
	fmt.Fprintln(out, "What is 5 - 7?", 5-7)
	fmt.Fprintln(out, "Oh, that's why it's false.")
	fmt.Fprintln(out, "How about some more.")
	fmt.Fprintln(out, "Is it greater?", 5 > -2)
	fmt.Fprintln(out, "Is it greater or equal?", 5 >= -2)
	fmt.Fprintln(out, "Is it less or equal?", 5 <= -2)
}

func main() {
	PrintMath(os.Stdout)
}
