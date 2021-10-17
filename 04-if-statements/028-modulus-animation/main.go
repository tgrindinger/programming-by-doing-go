package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func DoAnimation(times int, writer io.Writer) {
	for i := 0; i < times; i++ {
		if i % 22 == 0 {
			fmt.Fprint(writer, "*****           \r")
		} else if i % 22 == 1 {
			fmt.Fprint(writer, " *****          \r")
		} else if i % 22 == 2 {
			fmt.Fprint(writer, "  *****         \r")
		} else if i % 22 == 3 {
			fmt.Fprint(writer, "   *****        \r")
		} else if i % 22 == 4 {
			fmt.Fprint(writer, "    *****       \r")
		} else if i % 22 == 5 {
			fmt.Fprint(writer, "     *****      \r")
		} else if i % 22 == 6 {
			fmt.Fprint(writer, "      *****     \r")
		} else if i % 22 == 7 {
			fmt.Fprint(writer, "       *****    \r")
		} else if i % 22 == 8 {
			fmt.Fprint(writer, "        *****   \r")
		} else if i % 22 == 9 {
			fmt.Fprint(writer, "         *****  \r")
		} else if i % 22 == 10 {
			fmt.Fprint(writer, "          ***** \r")
		} else if i % 22 == 11 {
			fmt.Fprint(writer, "           *****\r")
		} else if i % 22 == 12 {
			fmt.Fprint(writer, "          ***** \r")
		} else if i % 22 == 13 {
			fmt.Fprint(writer, "         *****  \r")
		} else if i % 22 == 14 {
			fmt.Fprint(writer, "        *****   \r")
		} else if i % 22 == 15 {
			fmt.Fprint(writer, "       *****    \r")
		} else if i % 22 == 16 {
			fmt.Fprint(writer, "      *****     \r")
		} else if i % 22 == 17 {
			fmt.Fprint(writer, "     *****      \r")
		} else if i % 22 == 18 {
			fmt.Fprint(writer, "    *****       \r")
		} else if i % 22 == 19 {
			fmt.Fprint(writer, "   *****        \r")
		} else if i % 22 == 20 {
			fmt.Fprint(writer, "  *****         \r")
		} else if i % 22 == 21 {
			fmt.Fprint(writer, " *****          \r")
		}
		time.Sleep(200)
	}
}

func main() {
	DoAnimation(200, os.Stdout)
}
