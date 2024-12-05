package day04

import (
	"testing"

	"github.com/kabourneeak/aoc2024/internal/testutils"
	"github.com/stretchr/testify/assert"
)

// the example input from the puzzle
const exampleInput = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestRun_ShouldOutputResult(t *testing.T) {
	out := testutils.RunWithWriter(Run, exampleInput, t)

	const expectedOut = "Part 1 answer is 18\nPart 2 answer is 9\n"

	assert.Equal(t, expectedOut, out)
}
