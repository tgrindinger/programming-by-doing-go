package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func promptUser(scanner *bufio.Scanner, writer io.Writer, text string) string {
	fmt.Fprintln(writer, text)
	fmt.Fprint(writer, "> ")
	scanner.Scan()
	answer := scanner.Text()
	fmt.Fprintln(writer)
	return answer
}

func PlayGame(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprintln(writer, "WELCOME TO MITCHELL'S TINY ADVENTURE!")
	fmt.Fprintln(writer)
	answer := promptUser(scanner, writer, "You are in a creepy house!  Would you like to go \"upstairs\" or into the \"kitchen\"?")
	if answer == "kitchen" {
		answer = promptUser(scanner, writer, "There is a long countertop with dirty dishes everywhere.  Off to one side there is, as you'd expect, a refrigerator. You may open the \"refrigerator\" or look in a \"cabinet\".")
		if answer == "refrigerator" {
			answer = promptUser(scanner, writer, "Inside the refrigerator you see food and stuff.  It looks pretty nasty.  Would you like to eat some of the food? (\"yes\" or \"no\")")
			if answer == "yes" {
				fmt.Fprintln(writer, "You die of food poisoning.")
			} else {
				fmt.Fprintln(writer, "You die of starvation... eventually.")
			}
		} else {
			answer = promptUser(scanner, writer, "You see some food in the cabinets.  It looks pretty nasty.  Would you like to eat some of the food? (\"yes\" or \"no\")")
			if answer == "yes" {
				fmt.Fprintln(writer, "You die of food poisoning.")
			} else {
				fmt.Fprintln(writer, "You die of starvation... eventually.")
			}
		}
	} else {
		answer = promptUser(scanner, writer, "Upstairs you see a hallway.  At the end of the hallway is the master \"bedroom\".  There is also a \"bathroom\" off the hallway.  Where would you like to go?")
		if answer == "bedroom" {
			answer = promptUser(scanner, writer, "You are in a plush bedroom, with expensive-looking hardwood furniture.  The bed is unmade.  In the back of the room, the closet door is ajar.  Would you like to open the door? (\"yes\" or \"no\")")
			if answer == "yes" {
				fmt.Fprintln(writer, "A creepy clown jumps out and bonks you on the head.")
			} else {
				fmt.Fprintln(writer, "Well, then I guess you'll never know what was in there.  Thanks for playing, I'm tired of making nexted if statements.")
			}
		} else {
			answer = promptUser(scanner, writer, "You're in a grimy looking bathroom.  Would you like to take a bath? (\"yes\" or \"no\")")
			if answer == "yes" {
				fmt.Fprintln(writer, "You fail to notice the toaster precariously perched next to the tub until it's too late.")
			} else {
				fmt.Fprintln(writer, "Wallow in your filth, I guess.")
			}
		}
	}
}

func main() {
	PlayGame(os.Stdin, os.Stdout)
}
