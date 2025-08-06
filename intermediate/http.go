package main

import (
	"fmt"
	"net/url"
)

func Eval(err error, ret ...interface{}) []interface{} {
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return ret
}

func ParseUrl(myUrl string) (error, *url.URL) {
	parsedUrl, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("Failed to parse URL:", err)
		return err, nil
	}

	fmt.Printf("URL(%v):\n  Scheme: %v\n  Host: %v\n  Port: %v\n  Path: %v\n  RawQuery: %v\n  Fragment: %v\n",
		parsedUrl, parsedUrl.Scheme, parsedUrl.Host, parsedUrl.Port(), parsedUrl.Path, parsedUrl.RawQuery,
		parsedUrl.Fragment)

	params := parsedUrl.Query()
	fmt.Println(params)
	for key, value := range parsedUrl.Query() {
		fmt.Printf("    Key %s -> value %v  (len = %d)\n", key, value, len(value))
	}
	return nil, parsedUrl
}

func main() {
	//utl format: [schems://] [username@]host[:port][/path][?query][#fragment]
	rawUrl := "https://example.com:8080/path?query=param#fragment"
	Eval(ParseUrl(rawUrl))

	_, parsedUrl := ParseUrl("https://example.com/path?name=Jon&age=30") // Eval(ParseUrl("https://example.com/path?name=Jon&age=30"))
	params := parsedUrl.Query()
	fmt.Println(params)
	for key, value := range parsedUrl.Query() {
		fmt.Printf("Key %s -> value %v\n", key, value)
	}

	value := params["name"]
	fmt.Printf("Type of the value in parameter pair: %T\n", value)
	fmt.Println("params[\"name\"] =", params["name"])
	fmt.Println("params.Get(\"age\") =", params.Get("age"))
	fmt.Println("params.Get(\"XPEH\") =", params.Get("XPEH"))
	fmt.Println("params[\"XPEH\"] =", params["XPEH"])
	fmt.Println(params)

	//Build URL:
	baseUrl := &url.URL{
		Scheme: "https",
		Host:   "somewhere.far:666",
		//Port: 666,
		Path: "far/far/away",
		//RawQuery: "",
		Fragment: "1",
	}

	q := baseUrl.Query()
	q.Set("health", "good")
	q.Set("money", "a_lot")
	q.Add("money", "little_bit_more")
	baseUrl.RawQuery = q.Encode()
	fmt.Println(baseUrl)
	Eval(ParseUrl(baseUrl.String()))

	//building values for query:
	values := url.Values{}
	values.Add("name", "John")
	values.Add("city", "London")
	encodedQuery := values.Encode()
	fmt.Println("encodedQuery =", encodedQuery)

}
