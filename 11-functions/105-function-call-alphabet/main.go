package main

import (
	"fmt"
	"io"
	"os"
)

func a(writer io.Writer) {
	fmt.Fprint(writer, "Ant ")
}

func b() string {
	return "Banana "
}

func c(writer io.Writer, doit bool) {
	if doit {
		fmt.Fprint(writer, "Crocodile ")
	}
}

func d(writer io.Writer) {
	fmt.Fprint(writer, "Doggie ")
}

func e(writer io.Writer, howmany int) {
	s := "Elephant "
	x := 0
	for x < howmany {
		fmt.Fprintf(writer, "%c", s[x])
		x = x+1
	}
}

func f(writer io.Writer, word string) {
	fmt.Fprintf(writer, "%s ", word)
}

func g() string {
	return "Gorilla "
}

func h(writer io.Writer, reps int) {
	x := 0
	for x < reps {
		fmt.Fprint(writer, "Horseradish ")
		x = x+1
	}
}

func i(writer io.Writer, ignoredparameter int) {
	ignoredparameter = 32
	space := fmt.Sprintf("%c", ignoredparameter)
	fmt.Fprintf(writer, "%s%s", "Ice_cream", space)
}

func j(whichone int) string {
	if whichone == 1 {
		return "Jambalaya "
	} else if whichone == 2 {
		return "Juniper "
	} else {
		return "Jackrabbit "
	}
}

func k(writer io.Writer) {
	// the bird OR the fruit
	fmt.Fprint(writer, "Kiwi ")
}

func l(writer io.Writer, a, b bool) {
	if ( a && b ) {
		fmt.Fprint(writer, "Lettuce ")
	} else {
		fmt.Fprint(writer, "Lhasa_apso ")
	}
}

func m(writer io.Writer, a, b bool) {
	if a || b {
		fmt.Fprint(writer, "Mango ")
	} else {
		fmt.Fprint(writer, "Monkey! ")
	}
}

func n(writer io.Writer, word string) {
	fmt.Fprintf(writer, "%s ", word)
}

func o(writer io.Writer) int {
	fmt.Fprint(writer, "Orangutan ")
	return 1	// just for kicks; the return value doesn't mean anything
}

func p() string {
	return "Parrot "
}

func q(writer io.Writer) {
	fmt.Fprint(writer, "Quail ")
}

func r(first bool) string {
	if first {
		return "Rabbit "
	} else {
		return "Radish "
	}
}

func s(writer io.Writer, howmany int) {
	s := "Snake "
	x := 0
	for x < howmany {
		fmt.Fprintf(writer, "%c", s[x])
		x = x+1
	}
}

func t(writer io.Writer, word string) {
	fmt.Fprintf(writer, "%s ", word)
}

func u() string {
	return "Ugli_fruit "
}

func v(writer io.Writer, reps int) {
	x := 0
	for x < reps {
		fmt.Fprint(writer, "Valentine_candy ")
		x = x+1
	}
}

func w(writer io.Writer, ignoredparameter int) {
	ignoredparameter = 32
	space := fmt.Sprintf("%c", ignoredparameter)
	fmt.Fprintf(writer, "Walrus%s", space)
}

func x(whichone int) string {
	if whichone == 1 {
		return "X_files "
	} else {
		return "X_men "
	}
}

func y(writer io.Writer) {
	fmt.Fprint(writer, "Yams ")
}

func z(a, b bool) string {
	if a || b {
		return "Zanahorias "
	} else {
		return "Zebra "
	}
}

func PrintFunctionAlphabet(writer io.Writer) {
	a(writer)
	fmt.Fprint(writer, b())
	c(writer, true)
	d(writer)
	e(writer, 9)
	f(writer, "Frog")
	fmt.Fprint(writer, g())
	h(writer, 1)
	fmt.Fprintln(writer)

	i(writer, 5)
	fmt.Fprint(writer, j(3))
	k(writer)
	l(writer, true, false)
	m(writer, false, false)
	n(writer, "Narwhal")
	o(writer)
	fmt.Fprint(writer, p())
	q(writer)
	fmt.Fprintln(writer)

	fmt.Fprint(writer, r(true))
	s(writer, 6)
	t(writer, "Thyme")
	fmt.Fprint(writer, u())
	v(writer, 1)
	w(writer, 10)
	fmt.Fprint(writer, x(2))
	y(writer)
	fmt.Fprint(writer, z(false, false))
	fmt.Fprintln(writer)
}

func main() {
	PrintFunctionAlphabet(os.Stdout)
}
