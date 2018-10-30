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
		go func(link string, cachefolderpath string, index int) {
			defer wg.Done()
			getFileFromLink(link, cachefolderpath)
			messages <- fmt.Sprintf("getFileFromLink finished for %s, starting order: %d", link, index)
		}(link, cachefolderpath, i)
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

	// Convertir bytestream to string, autre technique?
	return fmt.Sprintf("%x", bs)
}

func getFileFromLink(link string, cachefolderpath string) {

	client := &http.Client{}

	req, _ := http.NewRequest("GET", link, nil)

	// Should user a user agent?
	// Detecter mimetype?
	// Creer un index?
	resp, err := client.Do(req)

	if err != nil {
		// todo : Pourquoi? Format de requete invalide?
		fmt.Printf("Failed to do request for link : %s\n", link)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// Note : will rewrite file if the file exist (refresh).
		filename := getSha1FileNameFromLink(link)
		// todo : extension?
		createFilePath := filepath.Join(cachefolderpath, filename)
		out, err := os.Create(createFilePath)
		if err != nil {
			panic(err)
		}

		defer out.Close()
		io.Copy(out, resp.Body)
		fmt.Printf("Downloaded link %s to : %s\n", link, createFilePath)
	}
}
