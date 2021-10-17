package main

import "fmt"

func Letter() string {
	return `+-------------------------------------------------------+
|                                                  #### |
|                                                  #### |
|                                                  #### |
|                                                       |
|                                                       |
|                        Bill Gates                     |
|                        1 Microsoft Way                |
|                        Redmond, WA 98104              |
|                                                       |
+-------------------------------------------------------+
`
}

func main() {
	fmt.Print(Letter())
}
