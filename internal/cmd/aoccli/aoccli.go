package aoccli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/kabourneeak/aoc2024/internal/day01"
	"github.com/kabourneeak/aoc2024/internal/day02"
	"github.com/kabourneeak/aoc2024/internal/day03"
	"github.com/kabourneeak/aoc2024/internal/day04"
	"github.com/kabourneeak/aoc2024/internal/days"
)

// cobra cli group for individual days
const dayGroupId = "day_group"

func RootCommand(name string) *cobra.Command {

	c := &cobra.Command{
		Use:   name,
		Short: "Run Advent of Code 2024 Puzzles",
	}

	c.AddCommand(
		EnvCommand(),
	)

	d := &cobra.Group{ID: dayGroupId, Title: "Days"}

	c.AddGroup(d)

	c.AddCommand(
		createDayCommand(1, day01.Run),
		createDayCommand(2, day02.Run),
		createDayCommand(3, day03.Run),
		createDayCommand(4, day04.Run),
	)

	return c
}

func createDayCommand(day int, run days.DayRunner) *cobra.Command {

	usage := fmt.Sprintf("day%d", day)
	short := fmt.Sprintf("Run Day %d Puzzles", day)

	c := &cobra.Command{
		Use:     usage,
		Short:   short,
		GroupID: dayGroupId,
		RunE: func(cmd *cobra.Command, args []string) error {
			env := GetEnv()
			input, err := GetPuzzleInput(env, day)

			if err != nil {
				return err
			}

			return run(input, cmd.OutOrStdout())
		},
	}

	return c
}
