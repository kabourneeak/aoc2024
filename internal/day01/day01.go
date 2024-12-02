package day01

import (
	"fmt"
	"io"
)

func Run(w io.Writer) {
	fmt.Fprintln(w, "Hello, World!")
}
