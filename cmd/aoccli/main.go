package main

import (
	"os"
	"path/filepath"

	"github.com/kabourneeak/aoc2024/internal/cmd/aoccli"
)

func main() {

	baseName := filepath.Base(os.Args[0])

	aoccli.RootCommand(baseName).Execute()
}
