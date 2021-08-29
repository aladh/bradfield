package main

import (
	"fmt"
	"os"

	"github.com/aladh/bradfield/0_prep/2_go/xkcd_search/index"
)

func main() {
	// Assume index already exists
	//err := index.Create()
	//if err != nil {
	//	log.Fatalln(err)
	//}

	results, err := index.Search(os.Args[1])
	if err != nil {
		fmt.Printf("error searching for %s: %s\n", os.Args[1], err)
	}

	fmt.Println(results)
}
