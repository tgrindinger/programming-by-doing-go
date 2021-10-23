package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func addKeychains(scanner *bufio.Scanner, writer io.Writer, keychains int) int {
	fmt.Fprintf(writer, "You have %d keychains. How many to add? ", keychains)
	var added int
	for ok := true; ok; ok = added < 0 {
		scanner.Scan()
		added, _ = strconv.Atoi(scanner.Text())
		if added < 0 {
			fmt.Fprintln(writer, "You cannot enter a negative number. Try again.")
		}
	}
	keychains += added
	fmt.Fprintf(writer, "You now have %d keychains.\n", keychains)
	return keychains
}

func removeKeychains(scanner *bufio.Scanner, writer io.Writer, keychains int) int {
	fmt.Fprintf(writer, "You have %d keychains. How many to remove? ", keychains)
	var removed int
	for ok := true; ok; ok = removed > keychains {
		scanner.Scan()
		removed, _ = strconv.Atoi(scanner.Text())
		if removed > keychains {
			fmt.Fprint(writer, "You don't have that many keychains. Try again: ")
		}
	}
	keychains -= removed
	fmt.Fprintf(writer, "You now have %d keychains.\n", keychains)
	return keychains
}

func viewOrder(writer io.Writer, keychains int, price, salesTax, shipping, perItem float64) {
	shippingCost := shipping + float64(keychains) * perItem
	subtotal := float64(keychains) * price + shippingCost
	tax := subtotal * salesTax
	fmt.Fprintf(writer, "You have %d keychains.\n", keychains)
	fmt.Fprintf(writer, "Keychains cost $%.2f each.\n", price)
	fmt.Fprintf(writer, "Shipping will be $%.2f.\n", shippingCost)
	fmt.Fprintf(writer, "Subtotal before tax: $%.2f.\n", subtotal)
	fmt.Fprintf(writer, "Tax will be $%.2f.\n", tax)
	fmt.Fprintf(writer, "Total cost is $%.2f.\n", subtotal + tax)
}

func checkout(scanner *bufio.Scanner, writer io.Writer, keychains int, price, salesTax, shipping, perItem float64) {
	fmt.Fprintln(writer, "CHECKOUT")
	fmt.Fprintln(writer)
	fmt.Fprint(writer, "What is your name? ")
	scanner.Scan()
	name := scanner.Text()
	viewOrder(writer, keychains, price, salesTax, shipping, perItem)
	fmt.Fprintf(writer, "Thanks for your order, %s!\n", name)
}

func SellKeychains(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprintln(writer, "Ye Olde Keychain Shoppe")
	keychains := 0
	price := 10.0
	salesTax := 0.0825
	shipping := 5.0
	perItem := 1.0
	var choice int
	for ok := true; ok; ok = choice != 4 {
		fmt.Fprintln(writer)
		fmt.Fprintln(writer, "1. Add Keychains to Order")
		fmt.Fprintln(writer, "2. Remove Keychains from Order")
		fmt.Fprintln(writer, "3. View Current Order")
		fmt.Fprintln(writer, "4. Checkout")
		fmt.Fprintln(writer)
		fmt.Fprint(writer, "Please enter your choice: ")
		scanner.Scan()
		fmt.Fprintln(writer)
		choice, _ = strconv.Atoi(scanner.Text())
		if choice == 1 {
			keychains = addKeychains(scanner, writer, keychains)
		} else if choice == 2 {
			keychains = removeKeychains(scanner, writer, keychains)
		} else if choice == 3 {
			viewOrder(writer, keychains, price, salesTax, shipping, perItem)
		} else if choice == 4 {
			checkout(scanner, writer, keychains, price, salesTax, shipping, perItem)
		} else {
			fmt.Fprintln(writer, "Please select a valid menu option!")
		}
	}
}

func main() {
	SellKeychains(os.Stdin, os.Stdout)
}
