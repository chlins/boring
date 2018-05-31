package main

import (
	"os"

	"github.com/chlins/boring/engine/bing"
)

func main() {
	if os.Args[1] != "" {
		w := bing.Translate(os.Args[1])
		if w != nil {
			bing.Print(w)
		}
	}
}
