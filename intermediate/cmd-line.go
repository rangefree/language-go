package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("os.Args:", os.Args)
	for i, e := range os.Args {
		fmt.Printf("osArgs[%2d]: %s\n", i, e)
	}

	//command line falds:
	var str string
	flag.StringVar(&str, "string", "DefaultValue", "Flag1 value to use for the run")

	var intFlag int
	flag.IntVar(&intFlag, "int", 0, "Flag1 value to use for the run")
	flag.Parse()

	fmt.Printf("Str: %s\n", str)
	fmt.Printf("IntFlag: %d\n", intFlag)

	//Subcommands:  go run .\cmd-line.go sub1 -processing=true -bytes=2024
	subcommand1 := flag.NewFlagSet("sub1", flag.ExitOnError)
	processing := subcommand1.Bool("processing", false, "Trigger to apply processing")
	bytes := subcommand1.Int("bytes", 1024, "Max byte length of result")

	subcommand2 := flag.NewFlagSet("sub2", flag.ExitOnError)
	language := subcommand2.String("language", "Go", "enter your language")

	if len(os.Args) < 2 {
		fmt.Println("Requires parameters.")
		os.Exit(1)
	}

	//--------------------------
	switch os.Args[1] {
	case "sub1":
		if err := subcommand1.Parse(os.Args[2:]); err != nil {
			fmt.Println("Failed to parse subcommand #1:", err)
			os.Exit(3)
		}
		fmt.Println("Subcommand 1\n  Processing", *processing, "\n  bytes", *bytes)
	case "sub2":
		subcommand2.Parse(os.Args[2:])
		fmt.Println("Subcommand 2\n  Language", *language)
	default:
		fmt.Printf("Unknown subcommand provided: %s\n", os.Args[1])
		os.Exit(2)
	}
}
