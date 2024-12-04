package day01

import (
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

func Run(input string, w io.Writer) error {

	model, err := parseInput(input)
	if err != nil {
		return err
	}

	part1, err := part1(model)
	if err != nil {
		return err
	}

	part2, err := part2(model)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "Part 1 answer is %d\n", part1)
	fmt.Fprintf(w, "Part 2 answer is %d\n", part2)

	return nil
}

type inputModel struct {
	ColA []int
	ColB []int
}

func parseInput(input string) (*inputModel, error) {
	lines := strings.Split(input, "\n")

	// break into two columns
	colA := make([]int, 0, len(lines))
	colB := make([]int, 0, len(lines))

	for _, line := range lines {
		a, b, err := splitLine(line)

		if err != nil {
			return nil, err
		}

		colA = append(colA, a)
		colB = append(colB, b)
	}

	m := &inputModel{
		ColA: colA,
		ColB: colB,
	}

	return m, nil
}

func part1(model *inputModel) (int, error) {
	// sort each column
	colA := slices.Clone(model.ColA)
	colB := slices.Clone(model.ColB)

	slices.Sort(colA)
	slices.Sort(colB)

	// pairwise sum the two columns
	var sum int

	for i := range colA {
		sum += abs(colA[i], colB[i])
	}

	return sum, nil
}

func part2(model *inputModel) (int, error) {
	colBfreq := make(map[int]int)

	for _, e := range model.ColB {
		if freq, ok := colBfreq[e]; ok {
			colBfreq[e] = freq + 1
		} else {
			colBfreq[e] = 1
		}
	}

	acc := 0

	for _, e := range model.ColA {
		if freq, ok := colBfreq[e]; ok {
			acc += e * freq
		}
	}

	return acc, nil
}

func abs(x, y int) int {
	if x < y {
		return y - x
	} else {
		return x - y
	}
}

func splitLine(line string) (int, int, error) {
	s := strings.Fields(line)
	if len(s) != 2 {
		return 0, 0, fmt.Errorf("line %s did not split into 2 parts", line)
	}

	a, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, 0, err
	}

	b, err := strconv.Atoi(s[1])
	if err != nil {
		return 0, 0, err
	}

	return a, b, err
}
