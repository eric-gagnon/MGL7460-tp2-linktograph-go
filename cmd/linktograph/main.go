// Package main is ...
package main

import (
	"fmt"
	"github.com/eric-gagnon/mgl7460-tp2-linktograph-go/pkg/link"
	"github.com/eric-gagnon/mgl7460-tp2-linktograph-go/pkg/pkg1"
	"github.com/eric-gagnon/mgl7460-tp2-linktograph-go/pkg/scraper"
	"path/filepath"
)

// main func is ...
func main() {
	fileOrganisationTesting()
	doProcess()
}

func fileOrganisationTesting() {
	fmt.Println("Starting the program.")
	somethingDone := pkg1.Dosomething("echo")
	somethingDone2 := pkg1.Dosomething2("echo")
	fmt.Println(somethingDone)
	fmt.Println(somethingDone2)
}

func doProcess() {
	fmt.Println("doProcess")

	absPath, _ := filepath.Abs(".")
	// todo: Input folder is hardcoded, could be a command-line parameter or config.
	readFilePath := filepath.Join(absPath, "input", "clean-links.txt") //"one-link-html.txt" // a-few-links.txt
	sourceLinks := link.GetUniqueLinksFromFile(readFilePath)

	scraper.ScrapFilesToCache(sourceLinks, filepath.Join(absPath, "cache", "web"))
}
