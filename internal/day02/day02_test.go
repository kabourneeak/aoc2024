package day02

import (
	"fmt"
	"testing"

	"github.com/kabourneeak/aoc2024/internal/testutils"
	"github.com/stretchr/testify/assert"
)

// the example input from the puzzle
const exampleInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestRun_ShouldOutputResult(t *testing.T) {
	out := testutils.RunWithWriter(Run, exampleInput, t)

	const expectedOut = "Part 1 answer is 2\nPart 2 answer is 4\n"

	assert.Equal(t, expectedOut, out)
}

func TestRun_ShouldPassSpecialCases(t *testing.T) {
	var tests = []struct {
		in        string
		wantPart1 int
		wantPart2 int
	}{
		// the first 64 must be skipped, but that is not clear while processing it
		// nor its predecessor.
		{"57 60 62 64 63 64 65", 0, 1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Report %s", tt.in)
		t.Run(testname, func(t *testing.T) {

			out := testutils.RunWithWriter(Run, tt.in, t)
			expected := fmt.Sprintf("Part 1 answer is %d\nPart 2 answer is %d\n", tt.wantPart1, tt.wantPart2)

			assert.Equal(t, out, expected)
		})
	}
}
