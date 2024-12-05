package day03

import (
	"testing"

	"github.com/kabourneeak/aoc2024/internal/testutils"
	"github.com/stretchr/testify/assert"
)

// the example input from the puzzle
const exampleInput = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func TestRun_ShouldOutputResult(t *testing.T) {
	out := testutils.RunWithWriter(Run, exampleInput, t)

	const expectedOut = "Part 1 answer is 161\nPart 2 answer is 48\n"

	assert.Equal(t, expectedOut, out)
}
