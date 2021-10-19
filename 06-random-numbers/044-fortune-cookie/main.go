package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func GiveFortune(random *rand.Rand, writer io.Writer) {
	fortune := random.Intn(6)
	if fortune == 0 {
		fmt.Fprintln(writer, "You will find happiness with a new love.");
	} else if fortune == 1 {
		fmt.Fprintln(writer, "You have good reason to be self-confident.");
	} else if fortune == 2 {
		fmt.Fprintln(writer, "Someone thinks the world of you.");
	} else if fortune == 3 {
		fmt.Fprintln(writer, "Your smile lights up someone else's day.");
	} else if fortune == 4 {
		fmt.Fprintln(writer, "Respect your elders. You could inherit a large sum of money.");
	} else if fortune == 5 {
		fmt.Fprintln(writer, "Face the truth with dignity.");
	}
	fmt.Fprintf(writer, "    %d - %d - %d - %d - %d - %d\n", 1 + random.Intn(54), 1 + random.Intn(54), 1 + random.Intn(54), 1 + random.Intn(54), 1 + random.Intn(54), 1 + random.Intn(54))
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	GiveFortune(random, os.Stdout)
}
