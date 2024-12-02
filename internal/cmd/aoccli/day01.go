package aoccli

import (
	"github.com/kabourneeak/aoc2024/internal/cmd/constants"
	"github.com/kabourneeak/aoc2024/internal/day01"
	"github.com/spf13/cobra"
)

func Day01Command() *cobra.Command {
	c := &cobra.Command{
		Use:     "day1",
		Short:   "Run Day 1 Puzzles",
		GroupID: constants.DayGroupId,
		RunE: func(cmd *cobra.Command, args []string) error {
			day01.Run(cmd.OutOrStdout())
			return nil
		},
	}

	return c
}
