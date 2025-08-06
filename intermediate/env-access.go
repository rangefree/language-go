package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

func main() {

	env := os.Environ()     // Get environment
	sort.Strings(env)       // sort slice
	for _, e := range env { //os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Printf("%s = %s\n", pair[0], pair[1])
		if pair[0] == "PATH" {

			//paths := pair[1]
			paths := strings.SplitN(pair[1], ":", -1)
			fmt.Println("Original:", paths)
			slices.Compact(paths)
			fmt.Println("Compacted:", paths)

			// fmt.Println("before:", paths)
			/*
				sort.Strings(paths)
				var newValue string
				for i:=0; i < len(paths) -1; i++ {
					newValue += paths[i]
					for i < len(paths)-1 && paths[i] == paths[i+1]  {
						i++
					}

				}

			*/
		}
	}

	fmt.Println("Get not existing env var 'WHATEVER':", os.Getenv("WHATEVER"))
}
