package day03

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
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
	Memory string
}

func parseInput(input string) (*inputModel, error) {
	model := &inputModel{Memory: input}
	return model, nil
}

func part1(input *inputModel) (int, error) {

	r, err := regexp.Compile(`mul\((\d\d?\d?),(\d\d?\d?)\)`)
	if err != nil {
		return 0, err
	}

	matches := r.FindAllStringSubmatch(input.Memory, -1)
	sum := 0

	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])

		sum += x * y
	}

	return sum, nil
}

func part2(input *inputModel) (int, error) {
	r, err := regexp.Compile(`(mul\((\d\d?\d?),(\d\d?\d?)\)|do\(\)|don't\(\))`)
	if err != nil {
		return 0, err
	}

	matches := r.FindAllStringSubmatch(input.Memory, -1)

	enabled := 1
	sum := 0

	for _, match := range matches {

		if match[0] == "do()" {
			enabled = 1
		} else if match[0] == "don't()" {
			enabled = 0
		} else {
			// note that the extra capture group sets the x and y back by one extra index position
			x, _ := strconv.Atoi(match[2])
			y, _ := strconv.Atoi(match[3])

			sum += x * y * enabled
		}
	}

	return sum, nil
}
