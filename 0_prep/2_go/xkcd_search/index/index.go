package index

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const baseURL = "https://xkcd.com"
const externalFilename = "info.0.json"
const notFound = 404

const storageDirectory = "comics"

func Create() error {
	for i := 2500; true; i++ {
		log.Printf("fetching comic %d\n", i)

		filename := fmt.Sprintf("%s/%d.json", storageDirectory, i)

		path, err := filepath.Abs(filename)
		if err != nil {
			return fmt.Errorf("error building file path for comic %d: %w", i, err)
		}

		if _, err := os.Stat(path); err == nil {
			log.Printf("skip fetching comic %d because a file already exists\n", i)
			continue
		} else if !errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("error checking existing file for comic %d: %w", i, err)
		}

		resp, err := http.Get(fmt.Sprintf("%s/%d/%s", baseURL, i, externalFilename))
		if err != nil {
			return fmt.Errorf("error making request for comic %d: %w", i, err)
		}

		if resp.StatusCode == notFound {
			log.Printf("comic %d not found - stopping index creation\n", i)
			break
		}

		f, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("error creating file for comic %d: %w", i, err)
		}

		_, err = io.Copy(f, resp.Body)
		if err != nil {
			return fmt.Errorf("error writing file for comic %d: %w", i, err)
		}

		_ = resp.Body.Close()
	}

	return nil
}
