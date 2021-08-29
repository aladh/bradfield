package fetcher

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const baseURL = "https://xkcd.com"
const filename = "info.0.json"
const notFound = 404

func Fetch(skip func(int) bool, store func(int, io.Reader) error) error {
	for i := 1; true; i++ {
		log.Printf("fetching comic %d\n", i)

		if skip(i) {
			log.Printf("skip fetching comic %d because a file already exists\n", i)
			continue
		}

		resp, err := http.Get(fmt.Sprintf("%s/%d/%s", baseURL, i, filename))
		if err != nil {
			return fmt.Errorf("error making request for comic %d: %w", i, err)
		}

		if resp.StatusCode == notFound && i != notFound { // comic 404 does not exist
			log.Printf("comic %d not found - stopped fetching\n", i)
			break
		}

		err = store(i, resp.Body)
		if err != nil {
			return err
		}

		_ = resp.Body.Close()
	}

	return nil
}
