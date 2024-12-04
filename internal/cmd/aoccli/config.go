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
			printConfig(cmd.OutOrStdout(), GetConfig())
			return nil
		},
	}

	return c
}

type CliConfig struct {
	BaseName string
	BaseDir  string
	WorkDir  string
	InputDir string
}

func GetConfig() CliConfig {
	baseName, _ := os.Executable()
	basePath := filepath.Dir(baseName)
	workDir, _ := os.Getwd()
	inputPath := filepath.Join(workDir, "inputs")

	return CliConfig{
		BaseName: baseName,
		BaseDir:  basePath,
		WorkDir:  workDir,
		InputDir: inputPath,
	}
}

func printConfig(w io.Writer, cfg CliConfig) {
	fmt.Fprintf(w, "Binary name: %s\n", cfg.BaseName)
	fmt.Fprintf(w, "Binary path: %s\n", cfg.BaseDir)
	fmt.Fprintf(w, "Current dir: %s\n", cfg.WorkDir)
	fmt.Fprintf(w, "Input dir: %s\n", cfg.InputDir)
}
