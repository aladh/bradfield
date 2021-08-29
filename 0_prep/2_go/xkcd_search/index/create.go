package index

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/aladh/bradfield/0_prep/2_go/xkcd_search/fetcher"
)

const storageDirectory = "comics"

func Create() error {
	err := fetcher.Fetch(skip, store)
	if err != nil {
		return err
	}

	return nil
}

func skip(comicNum int) bool {
	if _, err := os.Stat(path(comicNum)); err == nil {
		return true
	}

	return false
}

func store(comicNum int, comicReader io.Reader) error {
	f, err := os.Create(path(comicNum))
	if err != nil {
		return fmt.Errorf("error creating file for comic %d: %w", comicNum, err)
	}

	_, err = io.Copy(f, comicReader)
	if err != nil {
		return fmt.Errorf("error writing file for comic %d: %w", comicNum, err)
	}

	return nil
}

func path(comicNum int) string {
	filename := fmt.Sprintf("%s/%d.json", storageDirectory, comicNum)
	path, _ := filepath.Abs(filename)

	return path
}
