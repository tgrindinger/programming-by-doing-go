package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func randomMale(random *rand.Rand) string {
	names := []string{ "Rob", "Darius", "Cliff", "Dane", "Monty", "Reggie", "Gus",
			"Vernon", "Maynard", "Gavin", "Ward", "Stefan", "Winfred" }
	return names[random.Intn(len(names))]
}

func randomFemale(random *rand.Rand) string {
	names := []string{ "Laurene", "Kathy", "Yoko", "Matilda", "Georgette", "Greta",
			"Meg", "Ruby", "Rikki", "Suzanne", "Dorine", "Elaine", "Frederica" }
	return names[random.Intn(len(names))]
}

func randomNoun(random *rand.Rand) string {
	nouns := []string{ "moment", "joy", "license", "race", "inertia", "molecule",
			"helicopter", "sweat", "blame", "brush", "art", "song", "liberty" }
	return nouns[random.Intn(len(nouns))]
}

func randomVerbPastTense(random *rand.Rand) string {
	verbs := []string{ "ground", "settled", "gave", "labeled", "restored", "washed",
			"sold", "arranged", "neglected", "accused", "confined" }
	return verbs[random.Intn(len(verbs))]
}

func randomAdjective(random *rand.Rand) string {
	adjectives := []string{ "overdue", "collapsed", "jealous", "topical", "naughty", "third", 
			"sharp", "electrical", "bogus", "warm", "cryptic", "hopeful", "alleged" }
	return adjectives[random.Intn(len(adjectives))]
}

func randomAdverb(random *rand.Rand) string {
	adverbs := []string{ "endlessly", "wholly", "mentally", "brightly", "willingly",
			"sweetly", "legitimately", "behind", "instead", "within" }
	return adverbs[random.Intn(len(adverbs))]
}

func randomColor(random *rand.Rand) string {
	colors := []string{ "red", "orange", "yellow", "blue", "purple", "green",
			"white", "black", "pink", "brown", "grey", "hazel" }
	return colors[random.Intn(len(colors))]
}

func randomNumber(random *rand.Rand) string {
	numbers := []string{ "one", "two", "three", "four", "five", "six", "seven" }
	return numbers[random.Intn(len(numbers))]
}

func randomBodyPartPlural(random *rand.Rand) string {
	parts := []string{ "teeth", "legs", "arms", "ears", "eyes", "hands" }
	return parts[random.Intn(len(parts))]
}

func TellStory(random *rand.Rand, writer io.Writer) {
	male := randomMale(random)
	female := randomFemale(random)
	adj1 := randomAdjective(random)
	adj2 := randomAdjective(random)
	verb1 := randomVerbPastTense(random)
	verb2 := randomVerbPastTense(random)
	noun1 := randomNoun(random)
	noun2 := randomNoun(random)
	adv := randomAdverb(random)
	number := randomNumber(random)
	color := randomColor(random)
	color2 := randomColor(random)
	bodypart := randomBodyPartPlural(random)

	fmt.Fprintf(writer, "One afternoon %s and %s were walking down a(n) %s trail,\n", male, female, adj1)
	fmt.Fprintf(writer, "looking for kindling for their campfire. The trees were %s and\n", adj2)
	fmt.Fprintf(writer, "%s, and there were colorful wildflowers all around. %s and %s\n", color2, male, female)
	fmt.Fprintf(writer, "began to pick the wildflowers, and after a while, they %s so far\n", verb1)
	fmt.Fprintf(writer, "that they had wandered away from the trail.\n\n")

	fmt.Fprintf(writer, "It started to get dark. %s began to get worried, but %s seemed\n", male, female)
	fmt.Fprintf(writer, "excited to have an adventure. \"Look!\" %s said. \"Do you see that\n", female)
	fmt.Fprintf(writer, "%s? It looks like a house!\"\n\n", noun1)

	fmt.Fprintf(writer, "\"We're saved!\" cried %s, who was relieved.\n\n", male)

	fmt.Fprintf(writer, "Once they got closer, %s felt very uneasy again. It didn't look like\n", male)
	fmt.Fprintf(writer, "the cozy little cottage %s had been imagining, but rather a big, spooky\n", male)
	fmt.Fprintf(writer, "tower! It was about %s feet tall, and it was covered with %s ivy\n", number, color)
	fmt.Fprintf(writer, "and moss. It was the creepiest thing %s had ever seen!\n\n", male)

	fmt.Fprintf(writer, "%s said, \"%s, let's keep walking! There's no way I'm going into\n", male, female)
	fmt.Fprintf(writer, "that tower! It looks haunted!\"\n\n")

	fmt.Fprintf(writer, "\"Don't be such a(n) %s! We're going in. I think it looks %s\n", noun2, adv)
	fmt.Fprintf(writer, "un-haunted!\" said %s.\n\n", female)

	fmt.Fprintf(writer, "%s was so scared that he could not open his eyes. He felt his %s\n", male, bodypart)
	fmt.Fprintf(writer, "chatter as %s opened the door. All of a sudden, %s felt that he was\n", female, male)
	fmt.Fprintf(writer, "not alone. He opened his eyes, prepared to see the worst. But instead, he\n")
	fmt.Fprintf(writer, "saw all his friends and family inside the haunted tower! \"Surprise! Happy\n")
	fmt.Fprintf(writer, "birthday, %s!\" they all %s.\n", male, verb2)
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))
	TellStory(random, os.Stdout)
}
