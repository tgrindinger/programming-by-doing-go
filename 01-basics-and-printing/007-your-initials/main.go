package main

import "fmt"

func Initials() string {
	return `TTTTT JJJJJ  GGG
  T     J   G   G
  T     J   G
  T     J   GGGGG
  T   J J   G   G
  T   J J   G   G
  T    JJ    GGG
`
}

func main() {
	fmt.Print(Initials())
}
