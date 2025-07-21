package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("Hello, Base64 Encoding!")

	//encode:
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

	//decode:
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decoded:", string(decoded))

	//NOTE: avoid '/' and '+' if you want safe URL
	unsafeSlice := []byte("Using '/' and '+' in URL encoding is bad.")
	safeUrlEncoded := base64.StdEncoding.EncodeToString(unsafeSlice)
	fmt.Println("safe encoded:", safeUrlEncoded)
	safeDecoded, err := base64.StdEncoding.DecodeString(safeUrlEncoded)
	fmt.Println("safe decoded:", string(safeDecoded))
	//NOTE: use base64.URLEncoding for URLs

}
