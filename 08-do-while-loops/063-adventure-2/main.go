package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func DoAdventure(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprintln(writer, "MITCHELL'S TINY ADVENTURE 2!")
	status := "downstairs"
	for ok := true; ok; ok = status != "done" {
		fmt.Fprintln(writer)
		if status == "downstairs" {
			fmt.Fprintln(writer, "You are in a creepy house!  Would you like to go \"upstairs\" or into the \"kitchen\"?")
		} else if status == "kitchen" {
			fmt.Fprintln(writer, "There is a long countertop with dirty dishes everwhere.  Off to one side there is, as you'd expect, a refrigerator.  You may open the \"refrigerator\" or go \"back\".")
		} else if status == "upstairs" {
			fmt.Fprintln(writer, "Upstairs you see a hallway.  At the end of the hallway is the master \"bedroom\".  There is also a \"bathroom\" off the hallway.  Or, you can go back \"downstairs\". Where would you like to go?")
		} else if status == "refrigerator" {
			fmt.Fprintln(writer, "Inside the refrigerator you see food and stuff.  It looks pretty nasty.  Would you like to eat some of the food? (\"yes\" or \"no\")")
		} else if status == "bedroom" {
			fmt.Fprintln(writer, "There is, naturally, a bed.  It looks rather uninviting, but you could take a \"nap\" or go \"back\".")
		} else if status == "bathroom" {
			fmt.Fprintln(writer, "You feel pretty grimy.  You could take a \"bath\" or go \"back\".")
		}
		fmt.Fprint(writer, "> ")
		scanner.Scan()
		newStatus := scanner.Text()
		if status == "refrigerator" && newStatus == "yes" {
			fmt.Fprintln(writer, "The food is slimy and foul, but you manage to choke it down. Your stomach starts jumping like a frog in hot water.  You feel faint. Sliding to the floor, the darkness closes in.")
			fmt.Fprintln(writer)
			fmt.Fprintln(writer, "You have died.")
			status = "done"
		} else if newStatus == "nap" {
			fmt.Fprintln(writer, "The bed is lumpy and musty, but it serves its purpose.")
			fmt.Fprintln(writer)
			fmt.Fprintln(writer, "You have a pleasant nap.")
			status = "done"
		} else if newStatus == "bath" {
			fmt.Fprintln(writer, "You fail to notice the toaster precariously perched next to the tub until it's too late.")
			fmt.Fprintln(writer)
			fmt.Fprintln(writer, "You are crispy.")
			status = "done"
		} else if status == "kitchen" && newStatus == "back" {
			status = "downstairs"
		} else if status == "bedroom" && newStatus == "back" {
			status = "upstairs"
		} else if status == "bathroom" && newStatus == "back" {
			status = "upstairs"
		} else {
			status = newStatus
		}
	}
}

func main() {
	DoAdventure(os.Stdin, os.Stdout)
}
