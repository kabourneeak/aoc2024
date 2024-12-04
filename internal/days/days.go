package days

import "io"

type DayRunner func(input string, w io.Writer) error
