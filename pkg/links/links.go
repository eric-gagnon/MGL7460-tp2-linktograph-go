// Package implements utility routines to get a list of unique url links from a simple text file.
package links

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
)

// GetUniqueLinksFromFile get a list of unique url links from a simple text file.
func GetUniqueLinksFromFile(filepath string) []string {

	x := getLinksFromFile("clean-links.txt")
	x = removeDuplicates(x)

	return x
}

func getLinksFromFile(filename string) []string {

	links := []string{}

	absPath, _ := filepath.Abs(".")
	// todo: Input folder is hardcoded, could be a command-line parameter or config.
	readFilePath := filepath.Join(absPath, "input", filename)

	// See https://golang.org/pkg/bufio/#Scanner et https://gist.github.com/thedevsaddam/d7eff4608e2b41e2d8bc2b734183ede9
	fileHandle, _ := os.Open(readFilePath)
	defer fileHandle.Close()

	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {

		link := scanner.Text()
		// Minimal validation that the line is a valid url.
		// todo: Broken protocol is not filtered. Add test.
		_, err := url.ParseRequestURI(link)
		if err != nil {
			fmt.Printf("Link skipped, not parsed as a valid request URI: %s\n", link)
		} else {
			links = append(links, link)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	} else {
		fmt.Printf("Number of links found in file: %v\n", len(links))
	}

	return links
}

func removeDuplicates(s []string) []string {
	// Keeping order is not important.
	// No filtering on parameters

	filteredList := []string{}

	linkexist := map[string]bool{}

	for _, link := range s {
		if linkexist[link] {
			fmt.Printf("Duplicated link found: %s\n", link)
		} else {
			linkexist[link] = true
			filteredList = append(filteredList, link)
		}
	}

	fmt.Printf("Number of links without duplicates : %v\n", len(filteredList))

	return filteredList
}
