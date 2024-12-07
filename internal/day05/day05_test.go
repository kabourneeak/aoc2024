package day05

import (
	"testing"

	"github.com/kabourneeak/aoc2024/internal/testutils"
	"github.com/stretchr/testify/assert"
)

// the example input from the puzzle
const exampleInput = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func TestRun_ShouldOutputResult(t *testing.T) {
	out := testutils.RunWithWriter(Run, exampleInput, t)

	const expectedOut = "Part 1 answer is 143\nPart 2 answer is 123\n"

	assert.Equal(t, expectedOut, out)
}
