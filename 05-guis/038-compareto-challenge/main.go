package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func PrintComparison(writer io.Writer, first, second string) {
	fmt.Fprintf(writer, "Comparing \"%s\" with \"%s\" produces %d\n", first, second, strings.Compare(first, second))
}

func PrintComparisons(writer io.Writer) {
	PrintComparison(writer, "axe", "dog")
	PrintComparison(writer, "bug", "buzz")
	PrintComparison(writer, "cat", "frog")
	PrintComparison(writer, "deer", "giant")
	PrintComparison(writer, "money", "stuff")
	PrintComparison(writer, "beaver", "beaver")
	PrintComparison(writer, "custard", "custard")
	PrintComparison(writer, "bee", "artichoke")
	PrintComparison(writer, "frank", "betty")
	PrintComparison(writer, "zap", "apple")
	PrintComparison(writer, "tool", "charm")
	PrintComparison(writer, "things", "fun")
}

func main() {
	PrintComparisons(os.Stdout)
}