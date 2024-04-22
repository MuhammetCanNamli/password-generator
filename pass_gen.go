package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// Generate passwords with different lengths
	password1, err := generatePassword(12)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	password2, err := generatePassword(16)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Generated Password 1:", password1)
	fmt.Println("Generated Password 2:", password2)
}

func generatePassword(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("password length must be greater than 0")
	}

	// Defining password characters
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+"

	// Select random characters by looping through password length
	password := make([]byte, length)

	for i := range password {
		// Generate a random index within the charset length using crypto/rand
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[randomIndex.Int64()]
	}
	return string(password), nil
}
