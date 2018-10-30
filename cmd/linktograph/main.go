// Package main is ...
package main

import (
	"fmt"
	"github.com/eric-gagnon/mgl7460-tp2-linktograph-go/pkg/link"
	"github.com/eric-gagnon/mgl7460-tp2-linktograph-go/pkg/scraper"
	"path/filepath"
)

func main() {
	fmt.Println("Starting the processing.")
	doProcess()
	fmt.Println("Finished the processing.")
}

func doProcess() {
	fmt.Println("doProcess")

	absPath, _ := filepath.Abs(".")
	// todo: Input folder is hardcoded, could be a command-line parameter or config.
	readFilePath := filepath.Join(absPath, "input", "clean-links.txt") //"one-link-html.txt" // a-few-links.txt
	sourceLinks := link.GetUniqueLinksFromFile(readFilePath)

	scraper.ScrapFilesToCache(sourceLinks, filepath.Join(absPath, "cache", "web"))
}
