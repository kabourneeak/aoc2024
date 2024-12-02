package aoccli

import (
	"github.com/spf13/cobra"

	"github.com/kabourneeak/aoc2024/internal/cmd/constants"
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
		Day01Command(),
	)

	return c
}
