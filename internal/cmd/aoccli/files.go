package aoccli

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetPuzzleInput(cfg CliConfig, day int) (string, error) {
	file := fmt.Sprintf("day%d.txt", day)
	path := filepath.Join(cfg.InputDir, file)

	bytes, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	contents := string(bytes)

	return contents, nil
}
