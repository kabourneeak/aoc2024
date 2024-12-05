package day04

import (
	"fmt"
	"io"
	"slices"

	"github.com/kabourneeak/aoc2024/internal/days"
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
	Rows, Cols int
	Data       [][]rune
}

func parseInput(input string) (*inputModel, error) {

	lines := days.ToLines(input)

	if len(lines) == 0 {
		return nil, fmt.Errorf("input is empty")
	}

	// remove trailing new line
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	dataRows := make([][]rune, len(lines))

	for i, line := range lines {
		dataRows[i] = []rune(line)
	}

	model := &inputModel{
		Rows: len(lines),
		Cols: len(dataRows[0]),
		Data: dataRows,
	}

	return model, nil
}

const Xmas = "XMAS"

type Direction int64

const (
	N Direction = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

var Directions = []Direction{N, NE, E, SE, S, SW, W, NW}

var Diagonals = []Direction{NE, SE, SW, NW}

type Offset struct {
	X int
	Y int
}

func (d Direction) Offset() Offset {
	switch d {
	case N:
		return Offset{X: 0, Y: 1}
	case NE:
		return Offset{X: 1, Y: 1}
	case E:
		return Offset{X: 1, Y: 0}
	case SE:
		return Offset{X: 1, Y: -1}
	case S:
		return Offset{X: 0, Y: -1}
	case SW:
		return Offset{X: -1, Y: -1}
	case W:
		return Offset{X: -1, Y: 0}
	case NW:
		return Offset{X: -1, Y: 1}
	default:
		panic("Unhandled Case")
	}
}

func part1(input *inputModel) (int, error) {

	searchDir := func(input *inputModel, x int, y int, d Direction) int {
		offset := d.Offset()

		for i, l := range Xmas {
			posX := x + (offset.X * i)
			posY := y + (offset.Y * i)

			// went outside of data grid
			if posX < 0 || posX >= input.Cols || posY < 0 || posY >= input.Rows {
				return 0
			}

			// found wrong letter
			if input.Data[posY][posX] != l {
				return 0
			}
		}

		// made it!
		return 1
	}

	searchPos := func(input *inputModel, x int, y int) int {
		found := 0

		for _, d := range Directions {
			found += searchDir(input, x, y, d)
		}

		return found
	}

	search := func(input *inputModel) int {
		found := 0

		for r, row := range input.Data {
			for c := range row {
				found += searchPos(input, r, c)
			}
		}

		return found
	}

	found := search(input)

	return found, nil
}

func part2(input *inputModel) (int, error) {

	validCrosses := []string{
		"MMSS",
		"SMMS",
		"SSMM",
		"MSSM",
	}

	// the basic strategy here is to look for A's that the crosses may be centered around
	// then we collect the letters at the corners
	// then see if we have exactly the letters we expect
	searchPos := func(input *inputModel, x int, y int) int {
		// ensure we are inset far enough to collect the 4 corners of the cross
		if x < 1 || x >= input.Cols-1 || y < 1 || y >= input.Rows-1 {
			return 0
		}

		// see if we are around a potential cross center
		if input.Data[y][x] != 'A' {
			return 0
		}

		// collect our corners
		corners := make([]rune, len(Diagonals))

		for i, d := range Diagonals {
			corners[i] = input.Data[y+d.Offset().Y][x+d.Offset().X]
		}

		// see if it matches one of our templates
		if slices.Contains(validCrosses, string(corners)) {
			return 1
		}

		return 0
	}

	search := func(input *inputModel) int {
		found := 0

		for r, row := range input.Data {
			for c := range row {
				found += searchPos(input, r, c)
			}
		}

		return found
	}

	found := search(input)

	return found, nil
}
