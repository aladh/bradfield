package index

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Result struct {
	ComicNum   int
	Transcript string
}

func Search(searchTerm string) ([]string, error) {
	var results []string

	entries, err := os.ReadDir(storageDirectory)
	if err != nil {
		return results, fmt.Errorf("error reading directory: %w", err)
	}

	for _, entry := range entries {
		content, err := os.ReadFile(filepath.Join(storageDirectory, entry.Name()))
		if err != nil {
			return results, fmt.Errorf("error reading file %s: %w", entry.Name(), err)
		}

		if strings.Contains(strings.ToLower(string(content)), searchTerm) {
			results = append(results, entry.Name())
		}
	}

	return results, nil
}
