package day02

import (
	"fmt"
	"io"
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
	reports []report
}

type report struct {
	levels []int
}

func parseInput(input string) (*inputModel, error) {
	lines := strings.Split(input, "\n")

	parseLine := func(line string) ([]int, error) {
		fields := strings.Fields(line)
		levels := make([]int, len(fields))

		for i, s := range fields {
			level, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			levels[i] = level
		}

		return levels, nil
	}

	model := &inputModel{}

	for _, line := range lines {
		levels, err := parseLine(line)
		if err != nil {
			return nil, err
		}

		model.reports = append(model.reports, report{levels: levels})
	}

	return model, nil
}

func part1(input *inputModel) (int, error) {
	safeCount := 0
	for _, report := range input.reports {
		direction := calcDirection(report.levels)
		if safe := areLevelsSafe(report.levels, direction); safe {
			safeCount += 1
		}
	}

	return safeCount, nil
}

func part2(input *inputModel) (int, error) {

	evalSkips := func(levels []int) bool {
		direction := calcDirection(levels)

		for i := range levels {
			skipLevels := omitElement(levels, i)

			if areLevelsSafe(skipLevels, direction) {
				return true
			}
		}

		return false
	}

	safeCount := 0
	for _, report := range input.reports {
		if evalSkips(report.levels) {
			safeCount += 1
		}
	}

	return safeCount, nil
}

// calculate the overall trending direction of the report
// whether increasing or decreasing
func calcDirection(levels []int) int {
	direction := 0

	for i := 0; i < len(levels)-1; i += 1 {
		diff := levels[i+1] - levels[i]

		if diff > 0 {
			direction += 1
		} else if diff < 0 {
			direction -= 1
		}
	}

	if direction > 0 {
		return 1
	} else {
		return -1
	}
}

func areLevelsSafe(levels []int, direction int) bool {
	for i := 0; i < len(levels)-1; i += 1 {
		if !isLevelGood(levels[i+1], levels[i], direction) {
			return false
		}
	}

	return true
}

func isLevelGood(next, cur, direction int) bool {
	// diff is 0 -> delta == 0
	// diff is +ve, direction +ve -> delta > 0
	// diff is +ve, direction -ve -> delta < 0
	// diff is -ve, direction +ve -> delta < 0
	// diff is -ve, direction -ve -> delta > 0
	diff := next - cur
	delta := diff * direction

	return delta >= 1 && delta <= 3
}

func omitElement(s []int, i int) []int {

	// allocate new slice,
	// otherwise `s` itself will get modified.
	acc := make([]int, 0, len(s)-1)

	acc = append(acc, s[:i]...)
	acc = append(acc, s[i+1:]...)

	return acc
}
