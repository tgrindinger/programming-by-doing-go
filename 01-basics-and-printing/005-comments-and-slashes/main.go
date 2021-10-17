package main

import (
	"fmt"
	"io"
	"os"
)

func PrintCommentedLines(out io.Writer) {
	// A comment, this is so you can read your program later.
	// Anything after the // is ignored by Go
	fmt.Fprintln(out, "I could have code like this.") // and the comment after is ignored
	// You can also use a comment to "disable" or comment out a piece of code:
	// fmt.Fprintln("This won't run.")
	fmt.Fprintln(out, "This will run.")
}

func main() {
	PrintCommentedLines(os.Stdout)
}
