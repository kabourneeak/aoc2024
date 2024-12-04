package aoccli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func EnvCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "env",
		Short: "Display environment information for this tool",
		RunE: func(cmd *cobra.Command, args []string) error {
			printEnv(cmd.OutOrStdout(), GetEnv())
			return nil
		},
	}

	return c
}

type CliEnv struct {
	BaseName string
	BaseDir  string
	WorkDir  string
	InputDir string
}

func GetEnv() CliEnv {
	baseName, _ := os.Executable()
	basePath := filepath.Dir(baseName)
	workDir, _ := os.Getwd()
	inputPath := filepath.Join(workDir, "inputs")

	return CliEnv{
		BaseName: baseName,
		BaseDir:  basePath,
		WorkDir:  workDir,
		InputDir: inputPath,
	}
}

func printEnv(w io.Writer, env CliEnv) {
	fmt.Fprintf(w, "Binary name: %s\n", env.BaseName)
	fmt.Fprintf(w, "Binary path: %s\n", env.BaseDir)
	fmt.Fprintf(w, "Current dir: %s\n", env.WorkDir)
	fmt.Fprintf(w, "Input dir: %s\n", env.InputDir)
}
