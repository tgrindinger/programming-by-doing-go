package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func butterfly(writer io.Writer) {
	fmt.Fprintln(writer, "  .==-.                   .-==.     ")
	fmt.Fprintln(writer, "   \\()8`-._  `.   .'  _.-'8()/     ")
	fmt.Fprintln(writer, "   (88\"   ::.  \\./  .::   \"88)     ")
	fmt.Fprintln(writer, "    \\_.'`-::::.(#).::::-'`._/      ")
	fmt.Fprintln(writer, "      `._... .q(_)p. ..._.'         ")
	fmt.Fprintln(writer, "        \"\"-..-'|=|`-..-\"\"       ")
	fmt.Fprintln(writer, "        .\"\"' .'|=|`. `\"\".       ")
	fmt.Fprintln(writer, "      ,':8(o)./|=|\\.(o)8:`.        ")
	fmt.Fprintln(writer, "     (O :8 ::/ \\_/ \\:: 8: O)      ")
	fmt.Fprintln(writer, "      \\O `::/       \\::' O/       ")
	fmt.Fprintln(writer, "       \"\"--'         `--\"\"   hjw")
}

func elephant(writer io.Writer) {
	fmt.Fprintln(writer, "       _..--\"\"-.                  .-\"\"--.._ ")
	fmt.Fprintln(writer, "   _.-'         \\ __...----...__ /         '-._")
	fmt.Fprintln(writer, " .'      .:::...,'              ',...:::.      '.")
	fmt.Fprintln(writer, "(     .'``'''::;                  ;::'''``'.     )")
	fmt.Fprintln(writer, " \\             '-)              (-'             /")
	fmt.Fprintln(writer, "  \\             /                \\             /")
	fmt.Fprintln(writer, "   \\          .'.-.            .-.'.          /")
	fmt.Fprintln(writer, "    \\         | \\0|            |0/ |         /")
	fmt.Fprintln(writer, "    |          \\  |   .-==-.   |  /          |")
	fmt.Fprintln(writer, "     \\          `/`;          ;`\\`          /")
	fmt.Fprintln(writer, "      '.._      (_ |  .-==-.  | _)      _..'")
	fmt.Fprintln(writer, "          `\"`\"-`/ `/'        '\\` \\`-\"`\"`")
	fmt.Fprintln(writer, "               / /`;   .==.   ;`\\ \\")
	fmt.Fprintln(writer, "         .---./_/   \\  .==.  /   \\ \\")
	fmt.Fprintln(writer, "        / '.    `-.__)       |    `\"")
	fmt.Fprintln(writer, "       | =(`-.        '==.   ;")
	fmt.Fprintln(writer, " jgs    \\  '. `-.           /")
	fmt.Fprintln(writer, "         \\_:_)   `\"--.....-'")
}

func teddyBear(writer io.Writer) {
	fmt.Fprintln(writer, "            ___   .--. ")
	fmt.Fprintln(writer, "      .--.-\"   \"-' .- |")
	fmt.Fprintln(writer, "     / .-,`          .'")
	fmt.Fprintln(writer, "     \\   `           \\")
	fmt.Fprintln(writer, "      '.            ! \\")
	fmt.Fprintln(writer, "        |     !  .--.  |")
	fmt.Fprintln(writer, "        \\        '--'  /.____")
	fmt.Fprintln(writer, "       /`-.     \\__,'.'      `\\")
	fmt.Fprintln(writer, "    __/   \\`-.____.-' `\\      /")
	fmt.Fprintln(writer, "    | `---`'-'._/-`     \\----'    _")
	fmt.Fprintln(writer, "    |,-'`  /             |    _.-' `\\")
	fmt.Fprintln(writer, "   .'     /              |--'`     / |")
	fmt.Fprintln(writer, "  /      /\\              `         | |")
	fmt.Fprintln(writer, "  |   .\\/  \\      .--. __          \\ |")
	fmt.Fprintln(writer, "   '-'      '._       /  `\\         /")
	fmt.Fprintln(writer, "      jgs      `\\    '     |------'`")
	fmt.Fprintln(writer, "                 \\  |      |")
	fmt.Fprintln(writer, "                  \\        /")
	fmt.Fprintln(writer, "                   '._  _.'")
	fmt.Fprintln(writer, "                      ``")
}

func snake(writer io.Writer) {
	fmt.Fprintln(writer, "         ,,'6''-,.")
	fmt.Fprintln(writer, "        <====,.;;--.")
	fmt.Fprintln(writer, "        _`---===. \"\"\"==__")
	fmt.Fprintln(writer, "      //\"\"@@-\\===\\@@@@ \"\"\\\\")
	fmt.Fprintln(writer, "     |( @@@  |===|  @@@  ||")
	fmt.Fprintln(writer, "      \\\\ @@   |===|  @@  //")
	fmt.Fprintln(writer, "        \\\\ @@ |===|@@@ //")
	fmt.Fprintln(writer, "         \\\\  |===|  //")
	fmt.Fprintln(writer, "___________\\\\|===| //_____,----\"\"\"\"\"\"\"\"\"\"-----,_")
	fmt.Fprintln(writer, "  \"\"\"\"---,__`\\===`/ _________,---------,____    `,")
	fmt.Fprintln(writer, "             |==||                           `\\   \\")
	fmt.Fprintln(writer, "            |==| |          pb                 )   |")
	fmt.Fprintln(writer, "           |==| |       _____         ______,--'   '")
	fmt.Fprintln(writer, "           |=|  `----\"\"\"     `\"\"\"\"\"\"\"\"         _,-'")
	fmt.Fprintln(writer, "            `=\\     __,---\"\"\"-------------\"\"\"''")
	fmt.Fprintln(writer, "                \"\"\"\"		")
}

func PrintImage(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprintln(writer, "1. Butterfly ")
	fmt.Fprintln(writer, "2. Elephant  ")
	fmt.Fprintln(writer, "3. Teddy Bear")
	fmt.Fprintln(writer, "4. Snake     ")
	fmt.Fprintln(writer)
	fmt.Fprint(writer, "Which animal to draw? ")
	scanner.Scan()
	choice, _ := strconv.Atoi(scanner.Text())
	if choice == 1 {
		butterfly(writer)
	} else if choice == 2 {
		elephant(writer)
	} else if choice == 3 {
		teddyBear(writer)
	} else if choice == 4 {
		snake(writer)
	} else {
		fmt.Fprintln(writer, "Sorry, that wasn't one of the choices.")
	}
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Goodbye!")
}

func main() {
	PrintImage(os.Stdin, os.Stdout)
}
