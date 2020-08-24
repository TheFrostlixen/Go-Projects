package main

import (
	"os"
	"regexp"
	"fmt"
)

func main() {
	// get first argument passed to program
	orig_file := os.Args[1]
	
	// replace dots with spaces
	re1 := regexp.MustCompile(`\.`)
	filename := re1.ReplaceAllString(orig_file, " ")
	
	// perform filename formatting replace
	re2 := regexp.MustCompile(`(.*) (19|20)(\d{2}).* (.{3})$`)
	filename = re2.ReplaceAllString(filename, "$1 ($2$3).$4")
	
	// rename file
	os.Rename(orig_file, filename)
	fmt.Println(orig_file + " -> " + filename)
}
