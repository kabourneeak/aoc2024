package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/kabourneeak/aoc2024/internal/testutils"
)

func TestSplit_ShouldReturnBothSplits(t *testing.T) {
	a, b, err := splitLine("3   4")

	require.NoError(t, err)

	assert.Equal(t, 3, a)
	assert.Equal(t, 4, b)
}

func TestPart1_ShouldReturnAnswer(t *testing.T) {
	const input = `3   4
4   3
2   5
1   3
3   9
3   3`

	model, err := parseInput(input)
	require.NoError(t, err)

	result, err := part1(model)
	require.NoError(t, err)

	assert.Equal(t, 11, result)
}

func TestPart2_ShouldReturnAnswer(t *testing.T) {
	const input = `3   4
4   3
2   5
1   3
3   9
3   3`

	model, err := parseInput(input)
	require.NoError(t, err)

	result, err := part2(model)
	require.NoError(t, err)

	assert.Equal(t, 31, result)
}

func TestRun_ShouldOutputResult(t *testing.T) {
	const input = `3   4
	4   3
	2   5
	1   3
	3   9
	3   3`

	out := testutils.RunWithWriter(Run, input, t)

	const expectedOut = "Part 1 answer is 11\nPart 2 answer is 31\n"

	assert.Equal(t, expectedOut, out)
}
