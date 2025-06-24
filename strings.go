package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	str := "Hello go!"
	fmt.Println("01234567890")
	fmt.Println(str, "\tlength is", len(str))
	fmt.Println(str[0])

	fmt.Println("str[1:7] = ", str[1:7])
	fmt.Println("str[1:8] = ", str[1:8])
	fmt.Println("str[1:9] = ", str[1:9])

	str1 := "Hello"
	str2 := "World"
	fmt.Println((str1 + " " + str2))

	fmt.Printf("strconv.Itoa(255) = %T %s\n", strconv.Itoa(255), strconv.Itoa(255))
	m1 := "Hello \tworld" // c-style string
	m2 := `Hello \tworld` // raw string
	fmt.Printf("string (%s) has length %d\n", m1, len(m1))
	fmt.Printf("string (%s) has length %d\n", m2, len(m2))

	letters := "a,b,c,d,e"
	ll := strings.Split(letters, ",")
	fmt.Println(ll)
	fmt.Println(strings.Join(ll, "--"))
	fmt.Println(strings.Contains(letters, ",d"))
	fmt.Println(strings.Replace(letters, ",", ",__", 2))
	fmt.Println(strings.ToUpper(letters))
	fmt.Println(strings.Repeat("-", 20))

	fmt.Println("RegEx: ")
	str5 := "Hello 123 Go 11 1!"
	re := regexp.MustCompile(`\d+`)
	fmt.Println(re.FindAllString(str5, 5))

	fmt.Println("String Builder (fast!): ")
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World!")
	// res := builder.String()
	fmt.Println("Built string:", builder.String())
	builder.WriteRune(' ')
	builder.WriteString("Some crap...")
	fmt.Println("Built string:", builder.String())
}
