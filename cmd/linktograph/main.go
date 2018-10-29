// Package main is ...
package main

import (
	"fmt"
	"github.com/eric-gagnon/mgl7460-tp2-linktograph-go/pkg/links"
	"github.com/eric-gagnon/mgl7460-tp2-linktograph-go/pkg/pkg1"
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
	sourceLinks := links.GetUniqueLinksFromFile("")
	fmt.Println(len(sourceLinks))
}
