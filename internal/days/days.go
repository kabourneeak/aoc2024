package days

import "io"

// common interface for evaluating Day puzzles
type DayRunner func(input string, w io.Writer) error
