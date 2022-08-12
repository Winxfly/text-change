package main

import (
	"flag"
	"fmt"
	"log"
	"text-change/edit"
	"time"
)

func main() {
	t := time.Now()

	path := flag.String("path", "", "Path to directory")
	oldText := flag.String("old", "", "Old text")
	newText := flag.String("new", "", "New text")

	flag.Parse()

	if *path == "" {
		log.Fatal("[path is not specified]")
	}
	if *oldText == "" {
		log.Fatal("[old text is not specified]")
	}
	if *newText == "" {
		log.Fatal("[new text is not specified]")
	}

	edit.ReplaceTextInAllFilesDir(*path, *oldText, *newText)
	fmt.Printf("Execution completed. Time elapsed: %f seconds\n", time.Since(t).Seconds())
}
