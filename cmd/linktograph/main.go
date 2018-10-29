package main

import (
	"fmt"
	"github.com/eric-gagnon/mgl7460-tp2-linktograph-go/pkg/pkg1"
)

func main() {
	fmt.Println("Starting the program.")
	somethingDone := pkg1.Dosomething("echo")
	fmt.Println(somethingDone)
}
