// Package implements utility routines to get a list of unique url links from a simple text file.
package link

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"net/url"
	"os"
)

// GetUniqueLinksFromFile get a list of unique url links from a simple text file.
func GetUniqueLinksFromFile(filepath string) []string {

	x := getLinksFromFile(filepath)
	x = removeDuplicates(x)

	return x
}

func getLinksFromFile(filepath string) []string {

	links := []string{}

	// See https://golang.org/pkg/bufio/#Scanner et https://gist.github.com/thedevsaddam/d7eff4608e2b41e2d8bc2b734183ede9
	fileHandle, _ := os.Open(filepath)
	defer fileHandle.Close()

	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {

		l := scanner.Text()
		// Minimal validation that the line is a valid url.
		// todo: Broken protocol is not filtered. Add test.
		_, err := url.ParseRequestURI(l)
		if err != nil {
			fmt.Printf("Link skipped, not parsed as a valid request URI: %s\n", l)
		} else {
			links = append(links, l)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	} else {
		fmt.Printf("Number of links found in file: %v\n", len(links))
	}

	return links
}

func removeDuplicates(sourceLinks []string) []string {
	// Keeping order is not important.
	// No filtering on parameters

	filteredList := []string{}

	linkexist := map[string]bool{}

	for _, link := range sourceLinks {
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

func GetSha1FileNameForLink(link string) string {
	// Duplicate should be removed already but we want a simple way to generated a file name
	// that can be used as a access key later when accessing the cache.
	// Solution from : https://gobyexample.com/sha1-hashes
	h := sha1.New()
	h.Write([]byte(link))
	bs := h.Sum(nil)

	// todo : Convertir bytestream to string, autre technique?
	return fmt.Sprintf("%x", bs)
}