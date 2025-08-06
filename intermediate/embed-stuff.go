package main

import (
	"embed"
	"log"

	//	_ "embed" // using package for the "side effets". We want avoid go bitching about not used package: blank import
	"fmt"
	"io/fs"
)

// ERROR:
// //go:embed not-exist.txt
// var notExist string //must be of the suitamle type

//go:embed embed-stuff.txt
var embeddedStuff string //must be of the suitamle type

//go:embed basics
var basicsDir embed.FS

func main() {

	// Embeding files and directories directly into binaries:

	//ERROR: fmt.Println("Embedded content of not existing file:", notExist)

	fmt.Println("Embedded content:", embeddedStuff)
	content, err := basicsDir.ReadFile("basics/hello.go")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Embedded content from file basics/hello.go:\n%s\n", string(content))

	err = fs.WalkDir(basicsDir, "basics", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("%s %s\n", d.Type(), path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
