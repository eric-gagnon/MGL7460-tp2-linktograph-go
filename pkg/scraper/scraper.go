// Package implements utility routines to get files from the web.
package scraper

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type scrapedfile struct {
	url        string
	identifier string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ScrapFilesToCache(sourceLinks []string, cachefolderpath string) []scrapedfile {
	// Concurrency: https://nathanleclaire.com/blog/2014/02/15/how-to-wait-for-all-goroutines-to-finish-executing-before-continuing/
	messages := make(chan string)
	var wg sync.WaitGroup

	scrapedFiles := []scrapedfile{}

	wg.Add(len(sourceLinks))

	for i, link := range sourceLinks {

		filename := getSha1FileNameFromLink(link)
		cacheFilePath := filepath.Join(cachefolderpath, filename)

		go func(link string, cacheFilePath string, index int) {
			defer wg.Done()

			// https://golangcode.com/check-if-a-file-exists/
			if _, err := os.Stat(cacheFilePath); os.IsNotExist(err) {
				// todo : Ajouter retour erreur.
				downloadFileForLink(link, cacheFilePath)
				messages <- fmt.Sprintf("getFileFromLink finished for %s, starting order: %d", link, index)
			} else {
				messages <- fmt.Sprint("Skip downloadFileForLink, file already in cache.")
			}

		}(link, cacheFilePath, i)
	}

	go func() {
		for i := range messages {
			fmt.Println(i)
		}
	}()

	wg.Wait()

	return scrapedFiles
}

func getSha1FileNameFromLink(link string) string {
	// Duplicate should be removed already but we want a simple way to generated a file name
	// that can be used as a access key later when accessing the cache.
	// Solution from : https://gobyexample.com/sha1-hashes
	h := sha1.New()
	h.Write([]byte(link))
	bs := h.Sum(nil)

	// todo : Convertir bytestream to string, autre technique?
	return fmt.Sprintf("%x", bs)
}

func getFileForLink(link string, cachefolderpath string) {
}

func downloadFileForLink(link string, cacheFilePath string) {

	client := &http.Client{}

	// todo : handle the error?
	req, _ := http.NewRequest("GET", link, nil)

	resp, err := client.Do(req)

	if err != nil {
		// todo : add why.
		fmt.Printf("Skip download to cache, failed request. err: %v, link : %s,  \n", err, link)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// Note : will rewrite file if the file exist (refresh).

		out, err := os.Create(cacheFilePath)
		if err != nil {
			panic(err)
		}

		defer out.Close()
		io.Copy(out, resp.Body)
		fmt.Printf("Downloaded link %s to : %s\n", link, cacheFilePath)
	}
}
