package main

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	password := "password123"
	hash256 := sha256.Sum256([]byte(password)) // hash is a slice of bytes
	fmt.Println("sha256 hash:", hash256)
	fmt.Printf("sha256 hash in hex: %x\n", hash256)

	hash512 := sha512.Sum512([]byte(password)) // hash is a slice of bytes
	fmt.Printf("sha512 hash in hex: %x\n", hash512)

	// simulate preparation data to be stored in database: Base64(hashed salt+password) Base64(salt)
	salt, err := GenerateSalt() // random salt generation when creating password
	if err != nil {
		fmt.Println("Failed to create salt:", err)
		return
	}
	saltBase64 := base64.StdEncoding.EncodeToString(salt)    // Base64(salt)
	saltedPasswordHashBase64 := HashPassword(password, salt) //Base64(hashed salt+password)

	fmt.Printf("Base64 Salted Password: %x\nSalt: %x\nBase64 Salt: %x\n",
		saltedPasswordHashBase64, salt, saltBase64)

	// verify password:
	decodedSalt, _ := base64.StdEncoding.DecodeString(string(saltBase64))
	loginHash := HashPassword(password, decodedSalt)
	if loginHash == saltedPasswordHashBase64 {
		fmt.Printf("Verification: Password (%v) is correct.\n", password)
	} else {
		fmt.Println("Verification: Failed! Check credentials.")
	}
}

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func HashPassword(password string, salt []byte) string {
	saltedPass := append(salt, []byte(password)...)
	hash256 := sha256.Sum256(saltedPass)
	return base64.StdEncoding.EncodeToString(hash256[:])
}
