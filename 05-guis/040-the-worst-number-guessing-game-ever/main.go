package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func GuessNumber(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	answer := 4
	fmt.Fprintln(writer, "TEH WORST NUBMER GESSING GAME EVAR!!!!!!1!")
	fmt.Fprintln(writer)
	fmt.Fprint(writer, "I\"M THKING OF A NBR FROM 1-10.  TRY 2 GESS! ")
	scanner.Scan()
	guess, _ := strconv.Atoi(scanner.Text())
	fmt.Fprintln(writer)
	if guess == answer {
		fmt.Fprintf(writer, "LOL!!! U GOT IT!  I CANT BELEIVE U GESSED IT WAS %d!\n", answer)
	} else {
		fmt.Fprintf(writer, "W00T!  U SUX0R!!!  I PWN J00!!!  IT WAS %d!\n", answer)
	}
}

func main() {
	GuessNumber(os.Stdin, os.Stdout)
}
