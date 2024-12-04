package aoccli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/kabourneeak/aoc2024/internal/cmd/constants"
	"github.com/kabourneeak/aoc2024/internal/day01"
	"github.com/kabourneeak/aoc2024/internal/days"
)

func RootCommand(name string) *cobra.Command {

	c := &cobra.Command{
		Use:   name,
		Short: "Run Advent of Code 2024 Puzzles",
	}

	c.AddCommand(
		ConfigCommand(),
	)

	d := &cobra.Group{ID: constants.DayGroupId, Title: "Days"}

	c.AddGroup(d)

	c.AddCommand(
		createDayCommand(1, day01.Run),
	)

	return c
}

func createDayCommand(day int, run days.DayRunner) *cobra.Command {

	usage := fmt.Sprintf("day%d", day)
	short := fmt.Sprintf("Run Day %d Puzzles", day)

	c := &cobra.Command{
		Use:     usage,
		Short:   short,
		GroupID: constants.DayGroupId,
		RunE: func(cmd *cobra.Command, args []string) error {
			config := GetConfig()
			input, err := GetPuzzleInput(config, day)

			if err != nil {
				return err
			}

			return run(input, cmd.OutOrStdout())
		},
	}

	return c
}
