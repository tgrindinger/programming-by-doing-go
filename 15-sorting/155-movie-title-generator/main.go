package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func TellStory(random *rand.Rand, writer io.Writer) {
	nouns := []string{ "moment", "joy", "license", "race", "inertia", "molecule",
			"helicopter", "sweat", "blame", "brush", "art", "song", "liberty" }
	noun := nouns[random.Intn(len(nouns))]
	adjectives := []string{ "overdue", "collapsed", "jealous", "topical", "naughty", "third", 
			"sharp", "electrical", "bogus", "warm", "cryptic", "hopeful", "alleged" }
	adj := adjectives[random.Intn(len(adjectives))]

	fmt.Fprintln(writer, "Mitchell's Random Movie Title Generator")
	fmt.Fprintln(writer)
	fmt.Fprintf(writer, "Choosing randomly from %d adjectives and %d nouns (%d combinations).\n", len(adjectives), len(nouns), len(adjectives) * len(nouns))
	fmt.Fprintf(writer, "Your movie title is: %s %s\n", adj, noun)
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	TellStory(random, os.Stdout)
}
