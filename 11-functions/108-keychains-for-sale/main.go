package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func addKeychains(writer io.Writer) {
	fmt.Fprintln(writer, "ADD KEYCHAINS")
}

func removeKeychains(writer io.Writer) {
	fmt.Fprintln(writer, "REMOVE KEYCHAINS")
}

func viewOrder(writer io.Writer) {
	fmt.Fprintln(writer, "VIEW ORDER")
}

func checkout(writer io.Writer) {
	fmt.Fprintln(writer, "CHECKOUT")
}

func SellKeychains(reader io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(reader)
	fmt.Fprintln(writer, "Ye Olde Keychain Shoppe")
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
			addKeychains(writer)
		} else if choice == 2 {
			removeKeychains(writer)
		} else if choice == 3 {
			viewOrder(writer)
		} else if choice == 4 {
			checkout(writer)
		}
	}
}

func main() {
	SellKeychains(os.Stdin, os.Stdout)
}
