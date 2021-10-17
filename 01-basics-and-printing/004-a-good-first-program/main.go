package main

import "fmt"

func Lines() string {
	return `Hello World!
Hello Again
I like typing this.
This is fun.
Yay! Printing.
I'd much rather you 'not'.
I "said" do not touch this.`
}

func main() {
	fmt.Print(Lines())
}
