package main

import (
	"log"
	"os"
	"wget/internal/wget"
)

func main() {
	w := wget.New()

	if len(os.Args) < 2 {
		log.Fatal("please enter a link")
	}

	url := os.Args[1]

	if len(os.Args) == 3 {
		w.BasePath = os.Args[2]
	}

	w.VisitAndGet(url)
}
