package aoccli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func ConfigCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "config",
		Short: "Display configuration for this tool",
		RunE: func(cmd *cobra.Command, args []string) error {
			printConfig(cmd.OutOrStdout())
			return nil
		},
	}

	return c
}

func printConfig(w io.Writer) {
	baseName, _ := os.Executable()
	basePath := filepath.Dir(baseName)
	workDir, _ := os.Getwd()

	fmt.Fprintf(w, "Binary name: %s\n", baseName)
	fmt.Fprintf(w, "Binary path: %s\n", basePath)
	fmt.Fprintf(w, "Current dir: %s\n", workDir)
}
