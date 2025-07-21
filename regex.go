package main

import (
	"fmt"
	"regexp"
)

func IsEmail(email string) bool {
	//NOTE: "-" char has special meaning (range) if we have to use minus (-) as a symbol then it should be last
	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	match := re.MatchString(email) // name@domain.suffix
	fmt.Printf("String (%s) is email: %v\n", email, match)
	return match
}

func FindFirstSubmatches(s string, regexStr string) []string {
	re := regexp.MustCompile(regexStr)
	submatch := re.FindStringSubmatch(s)
	if submatch != nil {

	} else {
		fmt.Println("No matches found for regex (%s) in string (%s)", regexStr, s)
	}

	for i, v := range submatch {
		fmt.Printf("result[%d] = %s\n", i, v)
	}
	return submatch
}

func main() {
	// matching emails:
	IsEmail("user@email.com")
	IsEmail("user$@email.com")

	//Capture groups:
	//NOTE: result groups are defined in brackets (...)
	FindFirstSubmatches("2024-07-30 2025-06-25", `(\d{4})-(\d{2})-(\d{2})`)

	//Replacing:
	re := regexp.MustCompile(`[wold]`)
	fmt.Println("Hello World")
	fmt.Println(re.ReplaceAllString("Hello World", "."))

	// Flag syntax: ?[flag]
	// Flags:
	// i - case insensitive
	// m - multiline model
	// s - dot matches all
	re = regexp.MustCompile(`(?i)go`)
	text := "Golang is going great"
	fmt.Println("text matching:", re.MatchString(text))
	fmt.Println("text matching:", re.MatchString("Get better"))
}
