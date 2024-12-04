package testutils

import (
	"bytes"
	"io"
	"testing"

	"github.com/kabourneeak/aoc2024/internal/days"
)

func RunWithWriter(
	runFunc days.DayRunner,
	input string,
	t *testing.T,
) string {
	writer := bytes.NewBufferString("")

	// act
	runFunc(input, writer)

	// assert
	out, err := io.ReadAll(writer)
	if err != nil {
		t.Fatal(err)
	}

	return string(out)
}
