package utils

import (
	"os"
	"io"
	"fmt"
	"bufio"
	"path/filepath"
)

// ==============================
// CACHED ENCOURAGEMENTS
// ==============================
var cachedEncouragements []string

// ==============================
// LOAD ENCOURAGEMENTS
// ==============================
// LoadEncouragements loads the encouragements list from the config directory
// If missing, it recreates it from the embedded asset
func LoadEncouragements(_ string) ([]string, error) {
	if cachedEncouragements != nil {
		return cachedEncouragements, nil
	}

	configDir := filepath.Join(os.Getenv("HOME"), ".config", "cliwaifutamagotchi")
	if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed creating config dir: %w", err)
	}

	configPath := filepath.Join(configDir, "words-of-encouragement.txt")

	// Recreate default if missing
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := createEncFile(configPath); err != nil {
			return nil, err
		}
	}

	lines, err := readEncFile(configPath)
	if err != nil {
		return nil, err
	}

	cachedEncouragements = lines
	return lines, nil
}

// ==============================
// CREATE DEFAULT ENCOURAGEMENT FILE
// ==============================
func createEncFile(dst string) error {
	srcFile, err := ASSETSFS.Open("assets/words-of-encouragement.txt")
	if err != nil {
		return fmt.Errorf("missing embedded default encouragements: %w", err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create encouragements file: %w", err)
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return fmt.Errorf("failed to copy encouragements: %w", err)
	}

	return nil
}

// ==============================
// READ LINES FROM FILE
// ==============================
func readEncFile(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %w", path, err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed reading %s: %w", path, err)
	}

	return lines, nil
}
