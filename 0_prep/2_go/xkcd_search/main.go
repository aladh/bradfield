package main

import (
	"github.com/aladh/bradfield/0_prep/2_go/xkcd_search/index"
	"log"
)

func main() {
	err := index.Create()
	if err != nil {
		log.Fatalln(err)
	}
}
