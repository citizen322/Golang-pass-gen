package main

import {
	"fmt"
	"crypto/rand"
	"math/big"
}

const (
	lowercaseChars = "abcdefghijklmnopqrstuvwxyz"
	uppercaseChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitChars = "0123456789"
	specialChars = "!@#$%^&*()"
)

func generatePassword(length int) string {
	var passwordChars = make([]byte, length)
	var allChars = lowercaseChars + uppercaseChars + digitChars + specialChars
	allCharsLength := big.NewInt(int64(len(allChars)))

	for i := 0; i < length; ++i {
		randomIndex, _ := rand.Int(rand.Reader, allCharsLength)
		passwordChars[i] = allChars[randoxIndex.Int64]
	}

	return string(passwordChars)
}

func main() {
	var length int
	fmt.Print("Enter password length: ")
	fmt.Scan(&length)

	password := generatePassword(length)
	fmt.Println("Generated password: ", password)
}